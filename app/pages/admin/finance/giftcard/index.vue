<template>
  <div class="space-y-6">
    <!-- هدر صفحه -->
    <div class="bg-white rounded-lg shadow-sm p-6">
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-2xl font-bold text-gray-900">مدیریت جامع گیفت کارت</h1>
          <p class="text-gray-600 mt-1">مدیریت کامل گیفت کارت‌ها، تراکنش‌ها و گزارش‌ها</p>
        </div>
        <div class="flex items-center space-x-3 space-x-reverse">
          <button class="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors" @click="showCreateModal = true">
            <i class="fas fa-plus ml-2"></i>
            ایجاد گیفت کارت جدید
          </button>
          <button class="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors">
            <i class="fas fa-download ml-2"></i>
            گزارش
          </button>
        </div>
      </div>
    </div>

    <!-- کارت‌های آماری -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- کل کارت‌ها -->
      <div class="bg-white p-6 rounded-lg shadow-sm">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-sm font-medium text-gray-500">کل کارت‌ها</h3>
            <p class="text-2xl font-bold text-gray-900">۱,۲۵۰</p>
            <p class="text-sm text-green-600">۱۵+ <span class="text-gray-500">امروز</span></p>
          </div>
          <div class="w-12 h-12 bg-blue-500 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold">$</span>
          </div>
        </div>
      </div>
      
      <!-- کارت‌های فعال -->
      <div class="bg-white p-6 rounded-lg shadow-sm">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-sm font-medium text-gray-500">کارت‌های فعال</h3>
            <p class="text-2xl font-bold text-gray-900">۹۸۰</p>
            <p class="text-sm text-green-600">۷۸.۴% <span class="text-gray-500">نسبت کل</span></p>
          </div>
          <div class="w-12 h-12 bg-green-500 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold">✓</span>
          </div>
        </div>
      </div>
      
      <!-- در حال انقضا -->
      <div class="bg-white p-6 rounded-lg shadow-sm">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-sm font-medium text-gray-500">در حال انقضا</h3>
            <p class="text-2xl font-bold text-gray-900">۴۵</p>
            <p class="text-sm text-orange-600">۷ روز <span class="text-gray-500">آینده</span></p>
          </div>
          <div class="w-12 h-12 bg-yellow-500 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold">⏰</span>
          </div>
        </div>
      </div>
      
      <!-- ارزش کل -->
      <div class="bg-white p-6 rounded-lg shadow-sm">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-sm font-medium text-gray-500">ارزش کل</h3>
            <p class="text-2xl font-bold text-gray-900">۱۲۵,۰۰۰,۰۰۰ تومان</p>
            <p class="text-sm text-green-600">۱۲.۵%+ <span class="text-gray-500">ماه گذشته</span></p>
          </div>
          <div class="w-12 h-12 bg-red-500 rounded-lg flex items-center justify-center">
            <span class="text-white font-bold">$</span>
          </div>
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
      <div class="p-6">
        <!-- داشبورد -->
        <div v-if="activeTab === 'dashboard'" class="space-y-6">
          <GiftCardDashboard :gift-cards="giftCards" />
        </div>

        <!-- مدیریت کارت‌ها -->
        <div v-if="activeTab === 'cards'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">لیست کارت‌ها</h3>
              <GiftCardList :gift-cards="giftCards" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">جزئیات کارت</h3>
              <GiftCardDetails :card="selectedCard" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">مدیریت چرخه حیات</h3>
            <GiftCardLifecycleManagement :gift-cards="giftCards" />
          </div>
        </div>

        <!-- تراکنش‌ها -->
        <div v-if="activeTab === 'transactions'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">تراکنش‌ها</h3>
              <GiftCardTransactions :gift-cards="giftCards" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">فعالیت‌های اخیر</h3>
              <GiftCardRecentActivity :activities="[]" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">مدیریت بازگشت</h3>
            <GiftCardRefundManager :refunds="[]" />
          </div>
        </div>

        <!-- مدیریت کاربران -->
        <div v-if="activeTab === 'users'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت کاربران</h3>
              <GiftCardUserManager :users="[]" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت مشتریان</h3>
              <GiftCardCustomerManager :customers="[]" />
            </div>
          </div>
        </div>

        <!-- مدیریت محتوا -->
        <div v-if="activeTab === 'content'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت انواع کارت</h3>
              <GiftCardTypesManager :types="[]" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت قالب‌ها</h3>
              <GiftCardTemplateManager :templates="[]" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">مدیریت تم‌ها</h3>
            <GiftCardThemesManager :themes="[]" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">مدیریت پیام‌ها</h3>
            <GiftCardMessageManager :messages="[]" />
          </div>
        </div>

        <!-- گزارش‌ها و تحلیل -->
        <div v-if="activeTab === 'reports'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">گزارش‌های مالی</h3>
              <GiftCardFinancialReports :gift-cards="giftCards" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">گزارش‌های تحلیلی</h3>
              <GiftCardAnalyticsReports :gift-cards="giftCards" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">گزارش‌های کلی</h3>
            <GiftCardReports :gift-cards="giftCards" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">آمار کارت‌ها</h3>
            <GiftCardStats
:stats="{
              totalCards: 1250,
              activeCards: 980,
              usedCards: 200,
              expiredCards: 45,
              totalValue: 125000000,
              usedValue: 50000000,
              monthlyGrowth: 12.5,
              conversionRate: 78.4
            }" />
          </div>
        </div>

        <!-- مدیریت مالی -->
        <div v-if="activeTab === 'finance'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت مالی</h3>
              <GiftCardFinancialManagement :financial-data="{}" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت قیمت‌ها</h3>
              <GiftCardPricingManager :pricing-data="{}" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">قیمت‌گذاری</h3>
            <GiftCardPricing />
          </div>
        </div>

        <!-- جستجو و فیلتر -->
        <div v-if="activeTab === 'search'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">جستجوی سریع</h3>
              <GiftCardQuickSearch :gift-cards="giftCards" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">جستجوی پیشرفته</h3>
              <GiftCardAdvancedSearch :gift-cards="giftCards" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">جستجو و فیلتر</h3>
            <GiftCardSearchAndFilter :gift-cards="giftCards" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">فیلترها</h3>
            <GiftCardFilters :gift-cards="giftCards" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">فیلترهای ذخیره‌شده</h3>
            <GiftCardSavedFilters :current-filters="{}" />
          </div>
        </div>

        <!-- اعلان‌ها -->
        <div v-if="activeTab === 'notifications'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">مرکز اطلاع‌رسانی</h3>
              <GiftCardNotificationCenter :notifications="[]" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت اطلاع‌رسانی</h3>
              <GiftCardNotificationManager :notifications="[]" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">تنظیمات اطلاع‌رسانی</h3>
            <GiftCardNotificationSettings :settings="{}" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">تست اطلاع‌رسانی</h3>
            <GiftCardNotificationTest :test-data="{}" />
          </div>
        </div>

        <!-- تنظیمات -->
        <div v-if="activeTab === 'settings'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">تنظیمات عمومی</h3>
              <GiftCardGeneralSettings :settings="{}" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">تنظیمات امنیتی</h3>
              <GiftCardSecuritySettings :security-settings="{}" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">تنظیمات حریم خصوصی</h3>
            <GiftCardPrivacySettings :privacy-settings="{}" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">تنظیمات سیستمی</h3>
            <GiftCardSystemSettings :system-settings="{}" />
          </div>
        </div>

        <!-- ابزارها -->
        <div v-if="activeTab === 'tools'" class="space-y-6">
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <div>
              <h3 class="text-lg font-semibold mb-4">عملیات سریع</h3>
              <GiftCardQuickActions :gift-cards="giftCards" />
            </div>
            <div>
              <h3 class="text-lg font-semibold mb-4">مدیریت فیزیکی</h3>
              <GiftCardPhysicalManagement :physical-cards="[]" />
            </div>
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">بهینه‌سازی عملکرد</h3>
            <GiftCardPerformanceOptimization :performance-data="{}" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">امکانات پیشرفته</h3>
            <GiftCardAdvancedFeatures :features="[]" />
          </div>
          <div class="mt-6">
            <h3 class="text-lg font-semibold mb-4">راهنمای API</h3>
            <APIIntegrationGuide :api-data="{}" />
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- مودال ایجاد گیفت کارت -->
  <GiftCardCreateModal v-if="showCreateModal" @close="showCreateModal = false" />
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void

definePageMeta({ layout: 'admin-main' })

// تعریف interface ها
interface ApiResponse<T> {
  data?: T
  success?: boolean
  message?: string
}

interface GiftCard {
  id: string
  code: string
  value: number
  status: string
  createdAt: string
  expiresAt?: string
}

// ایمپورت کامپوننت‌های اصلی
import APIIntegrationGuide from './components/APIIntegrationGuide.vue';
import GiftCardAdvancedFeatures from './components/GiftCardAdvancedFeatures.vue';
import GiftCardAdvancedSearch from './components/GiftCardAdvancedSearch.vue';
import GiftCardAnalyticsReports from './components/GiftCardAnalyticsReports.vue';
import GiftCardCreateModal from './components/GiftCardCreateModal.vue';
import GiftCardCustomerManager from './components/GiftCardCustomerManager.vue';
import GiftCardDashboard from './components/GiftCardDashboard.vue';
import GiftCardDetails from './components/GiftCardDetails.vue';
import GiftCardFilters from './components/GiftCardFilters.vue';
import GiftCardFinancialManagement from './components/GiftCardFinancialManagement.vue';
import GiftCardFinancialReports from './components/GiftCardFinancialReports.vue';
import GiftCardGeneralSettings from './components/GiftCardGeneralSettings.vue';
import GiftCardLifecycleManagement from './components/GiftCardLifecycleManagement.vue';
import GiftCardList from './components/GiftCardList.vue';
import GiftCardMessageManager from './components/GiftCardMessageManager.vue';
import GiftCardNotificationCenter from './components/GiftCardNotificationCenter.vue';
import GiftCardNotificationManager from './components/GiftCardNotificationManager.vue';
import GiftCardNotificationSettings from './components/GiftCardNotificationSettings.vue';
import GiftCardNotificationTest from './components/GiftCardNotificationTest.vue';
import GiftCardPerformanceOptimization from './components/GiftCardPerformanceOptimization.vue';
import GiftCardPhysicalManagement from './components/GiftCardPhysicalManagement.vue';
import GiftCardPricing from './components/GiftCardPricing.vue';
import GiftCardPricingManager from './components/GiftCardPricingManager.vue';
import GiftCardPrivacySettings from './components/GiftCardPrivacySettings.vue';
import GiftCardQuickActions from './components/GiftCardQuickActions.vue';
import GiftCardQuickSearch from './components/GiftCardQuickSearch.vue';
import GiftCardRecentActivity from './components/GiftCardRecentActivity.vue';
import GiftCardRefundManager from './components/GiftCardRefundManager.vue';
import GiftCardReports from './components/GiftCardReports.vue';
import GiftCardSavedFilters from './components/GiftCardSavedFilters.vue';
import GiftCardSearchAndFilter from './components/GiftCardSearchAndFilter.vue';
import GiftCardSecuritySettings from './components/GiftCardSecuritySettings.vue';
import GiftCardStats from './components/GiftCardStats.vue';
import GiftCardSystemSettings from './components/GiftCardSystemSettings.vue';
import GiftCardTemplateManager from './components/GiftCardTemplateManager.vue';
import GiftCardThemesManager from './components/GiftCardThemesManager.vue';
import GiftCardTransactions from './components/GiftCardTransactions.vue';
import GiftCardTypesManager from './components/GiftCardTypesManager.vue';
import GiftCardUserManager from './components/GiftCardUserManager.vue';

// تعریف تب‌های اصلی - با توجه به ماهیت کانتینرها دسته‌بندی شده
const tabs = [
  // بخش اصلی
  { id: 'dashboard', title: 'داشبورد', icon: 'fas fa-tachometer-alt' },
  { id: 'cards', title: 'مدیریت کارت‌ها', icon: 'fas fa-credit-card' },
  { id: 'transactions', title: 'تراکنش‌ها', icon: 'fas fa-exchange-alt' },
  { id: 'users', title: 'مدیریت کاربران', icon: 'fas fa-users' },
  
  // بخش محتوا
  { id: 'content', title: 'مدیریت محتوا', icon: 'fas fa-edit' },
  
  // بخش گزارش‌گیری
  { id: 'reports', title: 'گزارش‌ها و تحلیل', icon: 'fas fa-chart-bar' },
  
  // بخش مالی
  { id: 'finance', title: 'مدیریت مالی', icon: 'fas fa-calculator' },
  
  // بخش جستجو
  { id: 'search', title: 'جستجو و فیلتر', icon: 'fas fa-search' },
  
  // بخش اعلان‌ها
  { id: 'notifications', title: 'اعلان‌ها', icon: 'fas fa-bell' },
  
  // بخش تنظیمات
  { id: 'settings', title: 'تنظیمات', icon: 'fas fa-cog' },
  
  // بخش ابزارها
  { id: 'tools', title: 'ابزارها', icon: 'fas fa-tools' }
]

// تب فعال
const activeTab = ref('dashboard')

// وضعیت نمایش مودال
const showCreateModal = ref(false)

const maxRetries = 10
const retryCount = ref(0)
const lastError = ref(null)
const loading = ref(false)
const giftCards = ref<GiftCard[]>([])

const fetchGiftCards = async (isManual = false) => {
  try {
    loading.value = true
    lastError.value = null
    // آدرس API را با توجه به بک‌اند خودت تنظیم کن
    const response: ApiResponse<GiftCard[]> = await $fetch('/api/giftcards')
    giftCards.value = response.data || []
    retryCount.value = 0 // موفقیت، شمارنده ریست شود
  } catch (error) {
    lastError.value = error
    if (!isManual && retryCount.value < maxRetries) {
      retryCount.value++
      setTimeout(() => fetchGiftCards(false), 1000)
    }
    // اگر دستی بود یا به سقف رسید، دیگر تلاش نکن
  } finally {
    loading.value = false
  }
}

const manualRefresh = () => {
  fetchGiftCards(true)
}

onMounted(() => {
  fetchGiftCards()
})

// کارت انتخاب شده برای نمایش جزئیات
const selectedCard = computed(() => giftCards.value[0])
</script>

<style scoped>
/* استایل‌های سفارشی برای صفحه گیفت کارت */
.space-x-reverse > * + * {
  margin-right: 0.75rem;
  margin-left: 0;
}

/* انیمیشن hover */
.transition-colors {
  transition: all 0.2s ease-in-out;
}
</style> 
