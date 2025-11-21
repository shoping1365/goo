<template>
  <div class="flex-1 p-6 space-y-8">
    <!-- تنظیمات امنیتی کوپن و کمپین -->
    <SecuritySettings />
    
    <!-- تولید خودکار کوپن -->
    <AutoCouponGenerator />
    
    <!-- لیست کوپن‌ها -->
    <CouponList 
      @show-add-form="showCouponForm = true; editingCoupon = null"
      @show-edit-form="editingCoupon = $event; showCouponForm = true"
      @show-details="selectedCoupon = $event; showCouponDetails = true"
    />
    
    <!-- لیست کمپین‌ها -->
    <CampaignList 
      @show-form="editingCampaign = $event; showCampaignForm = true"
    />
    
    <!-- مدیریت قوانین -->
    <RuleManagement />
    
    <!-- تخفیف‌های پویا -->
    <DynamicDiscounts />
    
    <!-- کمپین‌های شخصی‌سازی شده -->
    <PersonalizedCampaigns />
    
    <!-- گزارش‌های پیشرفته -->
    <AdvancedReports />
    
    <!-- قالب‌های پیشرفته -->
    <AdvancedTemplates />
    
    <!-- تنظیمات پیشرفته -->
    <AdvancedSettings />
    
    <!-- مدیریت اعلان‌ها -->
    <NotificationManagement />
    
    <!-- مدیریت محتوا -->
    <ContentManagement />
    
    <!-- تحلیل‌های کمپین -->
    <CampaignAnalytics />
    
    <!-- مدیریت بودجه -->
    <BudgetManagement />
    
    <!-- یکپارچه‌سازی سیستم -->
    <SystemIntegration />
    
    <!-- گزارش‌های تخفیف -->
    <DiscountReports />

    <!-- مودال فرم کوپن -->
    <CouponForm 
      v-if="showCouponForm"
      :coupon="editingCoupon"
      @save="handleCouponSave"
      @cancel="showCouponForm = false"
    />

    <!-- مودال جزئیات کوپن -->
    <CouponDetails 
      v-if="showCouponDetails"
      :coupon="selectedCoupon"
      @close="showCouponDetails = false"
      @edit="editingCoupon = $event; showCouponDetails = false; showCouponForm = true"
    />

    <!-- مودال فرم کمپین -->
    <CampaignForm 
      v-if="showCampaignForm"
      :campaign="editingCampaign"
      @save="handleCampaignSave"
      @cancel="showCampaignForm = false"
    />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'
// ایمپورت تمام کامپوننت‌ها
import AdvancedReports from './components/AdvancedReports.vue'
import AdvancedSettings from './components/AdvancedSettings.vue'
import AdvancedTemplates from './components/AdvancedTemplates.vue'
import AutoCouponGenerator from './components/AutoCouponGenerator.vue'
import BudgetManagement from './components/BudgetManagement.vue'
import CampaignAnalytics from './components/CampaignAnalytics.vue'
import CampaignForm from './components/CampaignForm.vue'
import CampaignList from './components/CampaignList.vue'
import ContentManagement from './components/ContentManagement.vue'
import CouponDetails from './components/CouponDetails.vue'
import CouponForm from './components/CouponForm.vue'
import CouponList from './components/CouponList.vue'
import DiscountReports from './components/DiscountReports.vue'
import DynamicDiscounts from './components/DynamicDiscounts.vue'
import NotificationManagement from './components/NotificationManagement.vue'
import PersonalizedCampaigns from './components/PersonalizedCampaigns.vue'
import RuleManagement from './components/RuleManagement.vue'
import SecuritySettings from './components/SecuritySettings.vue'
import SystemIntegration from './components/SystemIntegration.vue'

definePageMeta({ layout: 'admin-main' })

// State برای مودال‌ها
const showCouponForm = ref(false)
const showCouponDetails = ref(false)
const showCampaignForm = ref(false)
const editingCoupon = ref(null)
const editingCampaign = ref(null)
const selectedCoupon = ref(null)

interface Coupon {
  id?: number | string
  name?: string
  [key: string]: unknown
}

interface Campaign {
  id?: number | string
  name?: string
  [key: string]: unknown
}

// مدیریت ذخیره کوپن
function handleCouponSave(coupon: Coupon) {

  showCouponForm.value = false
  editingCoupon.value = null
  // TODO: فراخوانی API برای ذخیره کوپن
}

// مدیریت ذخیره کمپین
function handleCampaignSave(campaign: Campaign) {

  showCampaignForm.value = false
  editingCampaign.value = null
  // TODO: فراخوانی API برای ذخیره کمپین
}
</script>
