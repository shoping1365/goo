<template>
  <div class="shipping-methods-admin min-h-[100vh]">
    <!-- Header -->
    <div class="header-glass sticky top-0 z-20">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex items-center gap-6 py-7">
        <div class="flex items-center gap-6">
          <div class="header-icon flex items-center justify-center rounded-2xl bg-gradient-to-tr from-blue-400 via-purple-400 to-pink-400 shadow-lg w-16 h-16">
            <svg class="w-10 h-10 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V7"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h10M7 7V5a2 2 0 012-2h2a2 2 0 012 2v2"/></svg>
          </div>
          <div>
            <h1 class="text-3xl font-extrabold text-gray-900 tracking-tight">مدیریت روش‌های ارسال</h1>
            <p class="text-gray-500 text-base mt-1">مدیریت کامل تنظیمات، قیمت‌گذاری و ویژگی‌های روش‌های ارسال با تجربه کاربری حرفه‌ای</p>
          </div>
        </div>
        <div class="flex-1"></div>
        <div class="flex items-center gap-3">
          <button class="px-4 py-2 bg-gradient-to-r from-blue-500 to-purple-500 hover:from-blue-600 hover:to-purple-600 text-white rounded-xl text-sm font-bold shadow transition">راهنما</button>
          <button class="px-4 py-2 bg-gradient-to-r from-green-500 to-emerald-400 hover:from-green-600 hover:to-emerald-500 text-white rounded-xl text-sm font-bold shadow transition">پشتیبان‌گیری</button>
        </div>
      </div>
    </div>

    <!-- Navigation Tabs -->
    <div class="tab-glass sticky top-[90px] z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        <div class="relative">
          <button v-if="showScrollLeft" @click="scrollTabs('left')" class="tab-scroll-btn absolute left-0 top-1/2 -translate-y-1/2 z-10"><svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg></button>
          <nav ref="tabNav" class="flex flex-nowrap overflow-x-auto hide-scrollbar py-2 px-8" style="scroll-behavior:smooth;">
            <button
              v-for="(tab, idx) in tabs"
              :key="tab.name"
              @click="activeTab = idx"
              :class="[
                'tab-glass-btn flex items-center gap-2 px-7 py-2 rounded-xl font-bold text-base transition-all duration-200 whitespace-nowrap',
                activeTab === idx
                  ? tab.color + ' shadow-lg scale-105 ring-2 ring-offset-2 ring-blue-300 bg-opacity-90 text-white'
                  : 'bg-white/60 text-gray-700 hover:bg-blue-50 hover:text-blue-700',
                'min-w-[160px] sm:min-w-[180px] md:min-w-[200px]',
                activeTab === idx ? 'backdrop-blur-md' : ''
              ]"
            >
              <component :is="tab.icon" class="w-6 h-6" :class="activeTab === idx ? 'text-white' : tab.color.replace('bg-', 'text-')" />
              <span>{{ tab.label }}</span>
            </button>
          </nav>
          <button v-if="showScrollRight" @click="scrollTabs('right')" class="tab-scroll-btn absolute right-0 top-1/2 -translate-y-1/2 z-10"><svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg></button>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10">
      <transition name="fade-slide" mode="out-in">
        <div :key="activeTab" class="glass-card section-card bg-white/90 backdrop-blur-lg rounded-3xl shadow-2xl p-10 animate-fade-in">
          <div class="flex items-center gap-3 mb-6">
            <component :is="tabs[activeTab].icon" class="w-8 h-8" :class="tabs[activeTab].color.replace('bg-', 'text-')" />
            <h3 class="text-2xl font-extrabold text-gray-800">{{ tabs[activeTab].label }}</h3>
          </div>
          <p v-if="tabs[activeTab].desc" class="text-gray-500 mb-6">{{ tabs[activeTab].desc }}</p>
          <div class="section-card-inner">
            <component :is="tabs[activeTab].component" />
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
const { user, hasPermission } = useAuth()
// ایمپورت همه کامپوننت‌ها
import RequiredDocuments from './RequiredDocuments.vue'
import EmergencySettings from './EmergencySettings.vue'
import OrderValuePricing from './OrderValuePricing.vue'
import ClearDescriptions from './ClearDescriptions.vue'
import WeightLimits from './WeightLimits.vue'
import SystemSync from './SystemSync.vue'
import SecurityInsuranceSettings from './SecurityInsuranceSettings.vue'
import ActualDeliveryTime from './ActualDeliveryTime.vue'
import UsageStatistics from './UsageStatistics.vue'
import TimeSettings from './TimeSettings.vue'
import RegionalPricing from './RegionalPricing.vue'
import IssuesAndComplaints from './IssuesAndComplaints.vue'
import PackagingRequirements from './PackagingRequirements.vue'
import TimeBasedPricing from './TimeBasedPricing.vue'
import MethodComparison from './MethodComparison.vue'
import SelectionGuide from './SelectionGuide.vue'
import ShippingFAQ from './ShippingFAQ.vue'
import ShippingMethodImages from './ShippingMethodImages.vue'
import CountryRestrictions from './CountryRestrictions.vue'
import CustomsRegulations from './CustomsRegulations.vue'
import InternationalCurrencies from './InternationalCurrencies.vue'
import CompensationManagement from './CompensationManagement.vue'
import Compensation from './Compensation.vue'
import DelayNotifications from './DelayNotifications.vue'
import AlternativeShippingMethods from './AlternativeShippingMethods.vue'
import PersonalizationByHistory from './PersonalizationByHistory.vue'
import SuggestedShippingMethods from './SuggestedShippingMethods.vue'
import VipDiscounts from './VipDiscounts.vue'
import GroupVisibility from './GroupVisibility.vue'
import AutomaticNotifications from './AutomaticNotifications.vue'
import Webhooks from './Webhooks.vue'
import ExternalAPIs from './ExternalAPIs.vue'
import CustomerSatisfaction from './CustomerSatisfaction.vue'
import FinancialReports from './FinancialReports.vue'
import SpecialOrderConditions from './SpecialOrderConditions.vue'
import OrderTypeLimitations from './OrderTypeLimitations.vue'
import ItemCountLimits from './ItemCountLimits.vue'
import SensitiveProducts from './SensitiveProducts.vue'
import ProhibitedProducts from './ProhibitedProducts.vue'
import ProductSettings from './ProductSettings.vue'
import CoverageAreas from './CoverageAreas.vue'
import GeographicRange from './GeographicRange.vue'
import OrderSettings from './OrderSettings.vue'
import RegionalSettings from './RegionalSettings.vue'
import PricingSettings from './PricingSettings.vue'
import SpecialDiscountPricing from './SpecialDiscountPricing.vue'
import ItemCountPricing from './ItemCountPricing.vue'
import DistanceBasedPricing from './DistanceBasedPricing.vue'
import WeightBasedPricing from './WeightBasedPricing.vue'
import FlatRatePricing from './FlatRatePricing.vue'
import BaseInfoForm from './BaseInfoForm.vue'

// آیکون‌های رنگی برای هر تب
const IconInfo = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="12" r="10" stroke-width="2"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v-4m0-4h.01"/></svg>' }
const IconDoc = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7v10a2 2 0 002 2h6a2 2 0 002-2V7"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h10M7 7V5a2 2 0 012-2h2a2 2 0 012 2v2"/></svg>' }
const IconSetting = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3"/><circle cx="12" cy="12" r="10" stroke-width="2"/></svg>' }
const IconChart = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17v-6a2 2 0 012-2h2a2 2 0 012 2v6m4 0v-2a2 2 0 012-2h2a2 2 0 012 2v2"/></svg>' }
const IconBox = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><rect width="20" height="14" x="2" y="5" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 7l10 6 10-6"/></svg>' }
const IconList = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/></svg>' }
const IconStar = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><polygon points="12 2 15 8.5 22 9.3 17 14.1 18.2 21 12 17.8 5.8 21 7 14.1 2 9.3 9 8.5 12 2"/></svg>' }
const IconMoney = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="12" r="10" stroke-width="2"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3"/></svg>' }
const IconMap = { template: '<svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A2 2 0 013 15.382V5.618a2 2 0 011.553-1.894L9 2m6 0l5.447 2.724A2 2 0 0121 5.618v9.764a2 2 0 01-1.553 1.894L15 20m-6 0V2m6 18V2"/></svg>' }

const tabs = [
  { label: 'اطلاعات پایه', component: BaseInfoForm, name: 'BaseInfoForm', icon: IconInfo, color: 'bg-blue-500' },
  { label: 'مدارک مورد نیاز', component: RequiredDocuments, name: 'RequiredDocuments', icon: IconDoc, color: 'bg-purple-500' },
  { label: 'تنظیمات اضطراری', component: EmergencySettings, name: 'EmergencySettings', icon: IconSetting, color: 'bg-pink-500' },
  { label: 'قیمت‌گذاری بر اساس ارزش سفارش', component: OrderValuePricing, name: 'OrderValuePricing', icon: IconMoney, color: 'bg-orange-500' },
  { label: 'توضیحات شفاف', component: ClearDescriptions, name: 'ClearDescriptions', icon: IconInfo, color: 'bg-blue-400' },
  { label: 'محدودیت وزن', component: WeightLimits, name: 'WeightLimits', icon: IconBox, color: 'bg-teal-500' },
  { label: 'همگام‌سازی سیستم', component: SystemSync, name: 'SystemSync', icon: IconSetting, color: 'bg-indigo-500' },
  { label: 'تنظیمات امنیت و بیمه', component: SecurityInsuranceSettings, name: 'SecurityInsuranceSettings', icon: IconSetting, color: 'bg-green-500' },
  { label: 'زمان تحویل واقعی', component: ActualDeliveryTime, name: 'ActualDeliveryTime', icon: IconChart, color: 'bg-blue-600' },
  { label: 'آمار استفاده', component: UsageStatistics, name: 'UsageStatistics', icon: IconChart, color: 'bg-yellow-500' },
  { label: 'تنظیمات زمانی', component: TimeSettings, name: 'TimeSettings', icon: IconSetting, color: 'bg-pink-400' },
  { label: 'قیمت‌گذاری منطقه‌ای', component: RegionalPricing, name: 'RegionalPricing', icon: IconMap, color: 'bg-purple-400' },
  { label: 'مشکلات و شکایات', component: IssuesAndComplaints, name: 'IssuesAndComplaints', icon: IconList, color: 'bg-red-500' },
  { label: 'الزامات بسته‌بندی', component: PackagingRequirements, name: 'PackagingRequirements', icon: IconBox, color: 'bg-teal-400' },
  { label: 'قیمت‌گذاری زمانی', component: TimeBasedPricing, name: 'TimeBasedPricing', icon: IconMoney, color: 'bg-orange-400' },
  { label: 'مقایسه روش‌ها', component: MethodComparison, name: 'MethodComparison', icon: IconList, color: 'bg-blue-300' },
  { label: 'راهنمای انتخاب', component: SelectionGuide, name: 'SelectionGuide', icon: IconInfo, color: 'bg-blue-500' },
  { label: 'سوالات متداول', component: ShippingFAQ, name: 'ShippingFAQ', icon: IconInfo, color: 'bg-blue-400' },
  { label: 'تصاویر روش ارسال', component: ShippingMethodImages, name: 'ShippingMethodImages', icon: IconBox, color: 'bg-teal-300' },
  { label: 'محدودیت‌های کشوری', component: CountryRestrictions, name: 'CountryRestrictions', icon: IconMap, color: 'bg-purple-300' },
  { label: 'مقررات گمرکی', component: CustomsRegulations, name: 'CustomsRegulations', icon: IconDoc, color: 'bg-purple-500' },
  { label: 'ارزهای بین‌المللی', component: InternationalCurrencies, name: 'InternationalCurrencies', icon: IconMoney, color: 'bg-orange-500' },
  { label: 'مدیریت غرامت', component: CompensationManagement, name: 'CompensationManagement', icon: IconSetting, color: 'bg-green-400' },
  { label: 'غرامت', component: Compensation, name: 'Compensation', icon: IconSetting, color: 'bg-green-500' },
  { label: 'اعلان تاخیر', component: DelayNotifications, name: 'DelayNotifications', icon: IconSetting, color: 'bg-pink-500' },
  { label: 'روش‌های جایگزین', component: AlternativeShippingMethods, name: 'AlternativeShippingMethods', icon: IconList, color: 'bg-blue-400' },
  { label: 'شخصی‌سازی بر اساس تاریخچه', component: PersonalizationByHistory, name: 'PersonalizationByHistory', icon: IconStar, color: 'bg-yellow-400' },
  { label: 'پیشنهاد روش ارسال', component: SuggestedShippingMethods, name: 'SuggestedShippingMethods', icon: IconStar, color: 'bg-yellow-500' },
  { label: 'تخفیف ویژه VIP', component: VipDiscounts, name: 'VipDiscounts', icon: IconStar, color: 'bg-yellow-600' },
  { label: 'نمایش گروهی', component: GroupVisibility, name: 'GroupVisibility', icon: IconList, color: 'bg-blue-300' },
  { label: 'اعلان خودکار', component: AutomaticNotifications, name: 'AutomaticNotifications', icon: IconSetting, color: 'bg-green-400' },
  { label: 'وب‌هوک‌ها', component: Webhooks, name: 'Webhooks', icon: IconSetting, color: 'bg-indigo-400' },
  { label: 'APIهای خارجی', component: ExternalAPIs, name: 'ExternalAPIs', icon: IconSetting, color: 'bg-indigo-500' },
  { label: 'رضایت مشتری', component: CustomerSatisfaction, name: 'CustomerSatisfaction', icon: IconChart, color: 'bg-blue-500' },
  { label: 'گزارشات مالی', component: FinancialReports, name: 'FinancialReports', icon: IconMoney, color: 'bg-orange-500' },
  { label: 'شرایط ویژه سفارش', component: SpecialOrderConditions, name: 'SpecialOrderConditions', icon: IconSetting, color: 'bg-pink-400' },
  { label: 'محدودیت نوع سفارش', component: OrderTypeLimitations, name: 'OrderTypeLimitations', icon: IconList, color: 'bg-red-400' },
  { label: 'محدودیت تعداد آیتم', component: ItemCountLimits, name: 'ItemCountLimits', icon: IconList, color: 'bg-blue-400' },
  { label: 'محصولات حساس', component: SensitiveProducts, name: 'SensitiveProducts', icon: IconBox, color: 'bg-teal-400' },
  { label: 'محصولات ممنوعه', component: ProhibitedProducts, name: 'ProhibitedProducts', icon: IconBox, color: 'bg-teal-500' },
  { label: 'تنظیمات محصول', component: ProductSettings, name: 'ProductSettings', icon: IconSetting, color: 'bg-green-500' },
  { label: 'مناطق پوشش', component: CoverageAreas, name: 'CoverageAreas', icon: IconMap, color: 'bg-purple-400' },
  { label: 'محدوده جغرافیایی', component: GeographicRange, name: 'GeographicRange', icon: IconMap, color: 'bg-purple-500' },
  { label: 'تنظیمات سفارش', component: OrderSettings, name: 'OrderSettings', icon: IconSetting, color: 'bg-green-400' },
  { label: 'تنظیمات منطقه‌ای', component: RegionalSettings, name: 'RegionalSettings', icon: IconMap, color: 'bg-purple-300' },
  { label: 'تنظیمات قیمت‌گذاری', component: PricingSettings, name: 'PricingSettings', icon: IconMoney, color: 'bg-orange-400' },
  { label: 'قیمت‌گذاری تخفیف ویژه', component: SpecialDiscountPricing, name: 'SpecialDiscountPricing', icon: IconMoney, color: 'bg-orange-500' },
  { label: 'قیمت‌گذاری بر اساس تعداد آیتم', component: ItemCountPricing, name: 'ItemCountPricing', icon: IconMoney, color: 'bg-orange-400' },
  { label: 'قیمت‌گذاری بر اساس مسافت', component: DistanceBasedPricing, name: 'DistanceBasedPricing', icon: IconMoney, color: 'bg-orange-500' },
  { label: 'قیمت‌گذاری بر اساس وزن', component: WeightBasedPricing, name: 'WeightBasedPricing', icon: IconMoney, color: 'bg-orange-400' },
  { label: 'قیمت ثابت', component: FlatRatePricing, name: 'FlatRatePricing', icon: IconMoney, color: 'bg-orange-500' },
]

const activeTab = ref(0)
const tabNav = ref(null)
const showScrollLeft = ref(false)
const showScrollRight = ref(false)

function updateScrollButtons() {
  if (!tabNav.value) return
  const el = tabNav.value
  showScrollLeft.value = el.scrollLeft > 10
  showScrollRight.value = el.scrollWidth - el.clientWidth - el.scrollLeft > 10
}

function scrollTabs(dir) {
  if (!tabNav.value) return
  const el = tabNav.value
  const amount = 220
  if (dir === 'right') el.scrollLeft += amount
  else el.scrollLeft -= amount
  setTimeout(updateScrollButtons, 300)
}

onMounted(() => {
  if (tabNav.value) {
    tabNav.value.addEventListener('scroll', updateScrollButtons)
    updateScrollButtons()
  }
})
</script>

<style scoped>
.header-glass {
  background: rgba(255,255,255,0.85);
  backdrop-filter: blur(12px);
  border-bottom: 1.5px solid #e0e7ef;
  box-shadow: 0 4px 24px #bdbdf533;
}
.tab-glass {
  background: rgba(255,255,255,0.7);
  backdrop-filter: blur(8px);
  border-bottom: 1.5px solid #e0e7ef;
}
.tab-glass-btn {
  font-family: inherit;
  font-size: 1.08rem;
  font-weight: 700;
  border: none;
  outline: none;
  background: none;
  margin-bottom: -2px;
  box-shadow: none;
  transition: all 0.18s cubic-bezier(.4,0,.2,1);
  cursor: pointer;
  min-width: 120px;
  justify-content: center;
}
.tab-glass-btn:active {
  transform: scale(0.97);
}
.glass-card {
  box-shadow: 0 8px 40px #bdbdf544, 0 1.5px 8px #bdbdf522;
  background: rgba(255,255,255,0.85);
  backdrop-filter: blur(16px);
  border: 1.5px solid #e0e7ef;
}
.animate-fade-in {
  animation: fadeIn 0.5s;
}
.fade-slide-enter-active, .fade-slide-leave-active {
  transition: all 0.4s cubic-bezier(.4,0,.2,1);
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(30px) scale(0.98);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-20px) scale(0.98);
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: none; }
}
.hide-scrollbar::-webkit-scrollbar { display: none; }
.hide-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }
.tab-scroll-btn {
  background: rgba(255,255,255,0.8);
  border: none;
  outline: none;
  box-shadow: 0 2px 8px #bdbdf522;
  border-radius: 50%;
  width: 38px;
  height: 38px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 4px;
  cursor: pointer;
  transition: background 0.18s;
  z-index: 2;
}
.tab-scroll-btn:hover { background: #e3e8f7; }
.section-card {
  background: rgba(255,255,255,0.97);
  border-radius: 2rem;
  box-shadow: 0 8px 40px #bdbdf544, 0 1.5px 8px #bdbdf522;
  padding: 2.5rem 2rem;
  margin-bottom: 2rem;
  min-height: 400px;
}
.section-card-inner > * {
  margin-bottom: 1.5rem;
}
</style>







