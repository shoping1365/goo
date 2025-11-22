<template>
  <div class="flex-1 p-6 space-y-8">
    <!-- آمار کلی -->
    <ReferralDashboard />
    
    <!-- مدیریت کدهای ارجاع -->
    <ReferralCodes />
    
    <!-- قوانین پاداش‌دهی -->
    <RewardRules />
    
    <!-- لیست ارجاعات -->
    <ReferralList 
      @show-referral-details="selectedReferral = $event; showReferralDetails = true"
      @edit-referral="editingReferral = $event; showReferralForm = true"
    />
    
    <!-- گزارش‌ها و تحلیل -->
    <ReferralReports />
    
    <!-- تنظیمات برنامه -->
    <ReferralSettings />
    
    <!-- مدیریت درخواست‌های پاداش -->
    <RewardRequests />
    
    <!-- تاریخچه ارجاعات -->
    <ReferralHistory />
    
    <!-- مدیریت محتوا -->
    <ContentManagement />

    <!-- مودال جزئیات ارجاع -->
    <ReferralDetails 
      v-if="showReferralDetails"
      :referral="selectedReferral"
      @close="showReferralDetails = false"
      @edit="editingReferral = $event; showReferralDetails = false; showReferralForm = true"
    />

    <!-- مودال فرم ارجاع -->
    <ReferralForm 
      v-if="showReferralForm"
      :referral="editingReferral"
      @save="handleReferralSave"
      @cancel="showReferralForm = false"
    />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
import { ref } from 'vue'
// ایمپورت تمام کامپوننت‌ها
import ContentManagement from './components/ContentManagement.vue'
import ReferralCodes from './components/ReferralCodes.vue'
import ReferralDashboard from './components/ReferralDashboard.vue'
import ReferralDetails from './components/ReferralDetails.vue'
import ReferralForm from './components/ReferralForm.vue'
import ReferralHistory from './components/ReferralHistory.vue'
import ReferralList from './components/ReferralList.vue'
import ReferralReports from './components/ReferralReports.vue'
import ReferralSettings from './components/ReferralSettings.vue'
import RewardRequests from './components/RewardRequests.vue'
import RewardRules from './components/RewardRules.vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

// State برای مودال‌ها
const showReferralDetails = ref(false)
const showReferralForm = ref(false)
const selectedReferral = ref(null)
const editingReferral = ref(null)

interface Referral {
  id?: number
  code?: string
  userName?: string
  [key: string]: unknown
}

// مدیریت ذخیره ارجاع
function handleReferralSave(referral: Referral) {
  showReferralForm.value = false
  editingReferral.value = null
  // TODO: فراخوانی API برای ذخیره ارجاع
}
</script> 
