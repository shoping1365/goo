<template>
  <div class="flex-1 p-6 space-y-8">
    <!-- آمار کلی -->
    <LoyaltyDashboard />
    
    <!-- مدیریت سطوح وفاداری -->
    <LoyaltyLevels />
    
    <!-- قوانین امتیازدهی -->
    <PointRules />
    
    <!-- مدیریت پاداش‌ها -->
    <RewardsManagement />
    
    <!-- لیست اعضای وفاداری -->
    <LoyaltyMembers 
      @show-member-details="selectedMember = $event; showMemberDetails = true"
      @edit-member="editingMember = $event; showMemberForm = true"
    />
    
    <!-- گزارش‌ها و تحلیل -->
    <LoyaltyReports />
    
    <!-- تنظیمات برنامه -->
    <LoyaltySettings />
    
    <!-- مدیریت اعلان‌ها -->
    <LoyaltyNotifications />
    
    <!-- تاریخچه تراکنش‌ها -->
    <TransactionHistory />
    
    <!-- مدیریت محتوا -->
    <ContentManagement />

    <!-- مودال جزئیات عضو -->
    <MemberDetails 
      v-if="showMemberDetails"
      :member="selectedMember"
      @close="showMemberDetails = false"
      @edit="editingMember = $event; showMemberDetails = false; showMemberForm = true"
    />

    <!-- مودال فرم عضو -->
    <MemberForm 
      v-if="showMemberForm"
      :member="editingMember"
      @save="handleMemberSave"
      @cancel="showMemberForm = false"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
// ایمپورت تمام کامپوننت‌ها
import ContentManagement from './components/ContentManagement.vue'
import LoyaltyDashboard from './components/LoyaltyDashboard.vue'
import LoyaltyLevels from './components/LoyaltyLevels.vue'
import LoyaltyMembers from './components/LoyaltyMembers.vue'
import LoyaltyNotifications from './components/LoyaltyNotifications.vue'
import LoyaltyReports from './components/LoyaltyReports.vue'
import LoyaltySettings from './components/LoyaltySettings.vue'
import MemberDetails from './components/MemberDetails.vue'
import MemberForm from './components/MemberForm.vue'
import PointRules from './components/PointRules.vue'
import RewardsManagement from './components/RewardsManagement.vue'
import TransactionHistory from './components/TransactionHistory.vue'

definePageMeta({ layout: 'admin-main' })

// State برای مودال‌ها
const showMemberDetails = ref(false)
const showMemberForm = ref(false)
const selectedMember = ref(null)
const editingMember = ref(null)

interface Member {
  [key: string]: unknown
}

// مدیریت ذخیره عضو
function handleMemberSave(_member: Member) {
  showMemberForm.value = false
  editingMember.value = null
  // TODO: فراخوانی API برای ذخیره عضو
}
</script> 
