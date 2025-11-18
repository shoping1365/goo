<template>
  <div class="min-h-screen">
    <!-- Header Section -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4 space-x-reverse">
            <button 
              @click="toggleSidebar"
              class="p-2 rounded-md text-gray-500 hover:text-gray-700 hover:bg-gray-100 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
              </svg>
            </button>
            <div>
              <h1 class="text-xl font-bold text-gray-900">سیستم رتبه‌بندی کاربران</h1>
              <p class="mt-1 text-sm text-gray-500">مدیریت و تنظیم سیستم امتیازدهی و رتبه‌بندی کاربران</p>
            </div>
          </div>
          <div class="flex space-x-3 space-x-reverse">
            <button @click="activeTab = 'settings'" class="bg-blue-600 text-white px-3 py-2 rounded-md hover:bg-blue-700 transition-colors text-sm">
              تنظیمات سیستم
            </button>
            <button @click="exportData" class="bg-green-600 text-white px-3 py-2 rounded-md hover:bg-green-700 transition-colors text-sm">
              خروجی اکسل
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="flex">
      <!-- Mobile overlay -->
      <div 
        v-if="isSidebarOpen" 
        @click="toggleSidebar"
        class="fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden"
      ></div>
      
      <!-- Sidebar -->
      <div 
        :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full'"
        class="fixed inset-y-0 right-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out lg:relative lg:translate-x-0"
      >
        <div class="flex items-center justify-between p-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">منوها</h2>
          <button 
            @click="toggleSidebar"
            class="lg:hidden p-1 rounded-md text-gray-500 hover:text-gray-700 hover:bg-gray-100"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>
        
        <nav class="p-4 space-y-2">
          <!-- Main Menu Items -->
          <div class="space-y-1">
            <h3 class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">منوهای اصلی</h3>
            <button 
              @click="activeTab = 'dashboard'"
              :class="activeTab === 'dashboard' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              داشبورد
            </button>
            <button 
              @click="activeTab = 'users'"
              :class="activeTab === 'users' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              مدیریت کاربران
            </button>
            <button 
              @click="activeTab = 'scoring'"
              :class="activeTab === 'scoring' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              سیستم امتیازدهی
            </button>
            <button 
              @click="activeTab = 'manual'"
              :class="activeTab === 'manual' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              امتیازدهی دستی
            </button>
            <button 
              @click="activeTab = 'alerts'"
              :class="activeTab === 'alerts' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              سیستم هشدار
            </button>
            <button 
              @click="activeTab = 'benefits'"
              :class="activeTab === 'benefits' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              مزایای کاربران برتر
            </button>
          </div>

          <!-- Advanced Menu Items -->
          <div class="space-y-1 pt-4 border-t border-gray-200">
            <h3 class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">منوهای پیشرفته</h3>
            <button 
              @click="activeTab = 'reviews'"
              :class="activeTab === 'reviews' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              کیفیت نظرات
            </button>
            <button 
              @click="activeTab = 'referrals'"
              :class="activeTab === 'referrals' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              سیستم ارجاع
            </button>
            <button 
              @click="activeTab = 'purchases'"
              :class="activeTab === 'purchases' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              تداوم خرید
            </button>
            <button 
              @click="activeTab = 'returns'"
              :class="activeTab === 'returns' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              نرخ بازگشت
            </button>
            <button 
              @click="activeTab = 'account-age'"
              :class="activeTab === 'account-age' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              قدمت حساب
            </button>
          </div>

          <!-- System Menu Items -->
          <div class="space-y-1 pt-4 border-t border-gray-200">
            <h3 class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">سیستم</h3>
            <button 
              @click="activeTab = 'scenarios'"
              :class="activeTab === 'scenarios' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              سناریوها
            </button>
            <button 
              @click="activeTab = 'setup-test'"
              :class="activeTab === 'setup-test' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
            >
              راه‌اندازی و تست
            </button>
          </div>
        </nav>
      </div>

      <!-- Main Content -->
      <div class="flex-1 lg:ml-0">
        <div class="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
      <!-- Dashboard Tab -->
      <div v-if="activeTab === 'dashboard'">
        <ReportsAndAnalytics :users="users" @export-data="handleExportData" />
      </div>

      <!-- Users Management Tab -->
      <div v-if="activeTab === 'users'">
        <UserManagement 
          :users="users" 
          @view-details="handleViewUserDetails"
          @edit-user="handleEditUser"
          @toggle-status="handleToggleUserStatus"
        />
      </div>

      <!-- Auto Scoring System Tab -->
      <div v-if="activeTab === 'scoring'">
        <AutoScoringSystem 
          :initial-settings="scoringSettings"
          @save-settings="handleSaveScoringSettings"
        />
      </div>

      <!-- Manual Scoring Tab -->
      <div v-if="activeTab === 'manual'">
        <ManualScoring 
          :users="users"
          @submit-score="handleSubmitManualScore"
          @export-history="handleExportHistory"
        />
      </div>

      <!-- Alert System Tab -->
      <div v-if="activeTab === 'alerts'">
        <AlertSystem 
          :users="users"
          @send-alert="handleSendAlert"
          @send-bulk-alerts="handleSendBulkAlerts"
          @export-history="handleExportAlertHistory"
        />
      </div>

      <!-- Top User Benefits Tab -->
      <div v-if="activeTab === 'benefits'">
        <TopUserBenefits
          :users="users"
          @save-benefits="handleSaveBenefits"
          @export-data="handleExportBenefitsData"
        />
      </div>

      <!-- Review Quality System Tab -->
      <div v-if="activeTab === 'reviews'">
        <ReviewQualitySystem
          :reviews="reviews"
          @save-settings="handleSaveReviewSettings"
          @analyze-reviews="handleAnalyzeReviews"
          @export-report="handleExportReviewReport"
        />
      </div>

      <!-- Referral System Tab -->
      <div v-if="activeTab === 'referrals'">
        <ReferralSystem
          :referrals="referrals"
          @save-settings="handleSaveReferralSettings"
          @generate-codes="handleGenerateReferralCodes"
          @export-report="handleExportReferralReport"
        />
      </div>

      <!-- Purchase Continuity System Tab -->
      <div v-if="activeTab === 'purchases'">
        <PurchaseContinuitySystem
          :users="users"
          @save-settings="handleSaveContinuitySettings"
          @analyze-patterns="handleAnalyzePurchasePatterns"
          @export-report="handleExportContinuityReport"
        />
      </div>

      <!-- Return Rate System Tab -->
      <div v-if="activeTab === 'returns'">
        <ReturnRateSystem
          :users="users"
          @save-settings="handleSaveReturnSettings"
          @analyze-patterns="handleAnalyzeReturnPatterns"
          @export-report="handleExportReturnReport"
        />
      </div>

      <!-- Account Age System Tab -->
      <div v-if="activeTab === 'account-age'">
        <AccountAgeSystem
          :users="users"
          @save-settings="handleSaveAgeSettings"
          @analyze-ages="handleAnalyzeAccountAges"
          @export-report="handleExportAgeReport"
        />
      </div>

      <!-- System Settings Tab -->
      <div v-if="activeTab === 'settings'">
        <SystemSettings
          :initial-settings="systemSettings"
          @save-settings="handleSaveSystemSettings"
        />
      </div>

      <!-- Scenario Management System Tab -->
      <div v-if="activeTab === 'scenarios'">
        <ScenarioManagementSystem
          :scenarios="scenarios"
          @save-scenario="handleSaveScenario"
          @update-scenario="handleUpdateScenario"
          @delete-scenario="handleDeleteScenario"
          @export-scenarios="handleExportScenarios"
        />
      </div>

      <!-- Setup and Testing System Tab -->
      <div v-if="activeTab === 'setup-test'">
        <SetupAndTestingSystem
          @run-setup="handleRunSetup"
          @run-tests="handleRunTests"
          @view-results="handleViewTestResults"
        />
      </div>
    </div>
  </div>
</div>
</div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string }) => void
declare const navigateTo: (to: string) => Promise<void>
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import AccountAgeSystem from './components/AccountAgeSystem.vue';
import AlertSystem from './components/AlertSystem.vue';
import AutoScoringSystem from './components/AutoScoringSystem.vue';
import ManualScoring from './components/ManualScoring.vue';
import PurchaseContinuitySystem from './components/PurchaseContinuitySystem.vue';
import ReferralSystem from './components/ReferralSystem.vue';
import ReportsAndAnalytics from './components/ReportsAndAnalytics.vue';
import ReturnRateSystem from './components/ReturnRateSystem.vue';
import ReviewQualitySystem from './components/ReviewQualitySystem.vue';
import ScenarioManagementSystem from './components/ScenarioManagementSystem.vue';
import SetupAndTestingSystem from './components/SetupAndTestingSystem.vue';
import SystemSettings from './components/SystemSettings.vue';
import TopUserBenefits from './components/TopUserBenefits.vue';
import UserManagement from './components/UserManagement.vue';

definePageMeta({ layout: 'admin-main' })

// Reactive data
const activeTab = ref('dashboard')
const isSidebarOpen = ref(false)

// Toggle sidebar function
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

// Sample users data with updated structure
const users = ref([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/avatars/ali.jpg',
    score: 850,
    level: 'diamond',
    lastActivity: '2024-01-15T10:30:00Z',
    status: 'active',
    currentScore: 850
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/avatars/fateme.jpg',
    score: 720,
    level: 'platinum',
    lastActivity: '2024-01-14T15:45:00Z',
    status: 'active',
    currentScore: 720
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/avatars/mohammad.jpg',
    score: -75,
    level: 'bronze',
    lastActivity: '2024-01-13T09:20:00Z',
    status: 'blocked',
    currentScore: -75
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/avatars/zahra.jpg',
    score: 650,
    level: 'gold',
    lastActivity: '2024-01-15T14:10:00Z',
    status: 'active',
    currentScore: 650
  },
  {
    id: 5,
    name: 'حسین نوری',
    email: 'hossein@example.com',
    avatar: '/avatars/hossein.jpg',
    score: 150,
    level: 'silver',
    lastActivity: '2024-01-12T11:30:00Z',
    status: 'active',
    currentScore: 150
  }
])

// Sample reviews data
const reviews = ref([
  {
    id: 1,
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    productName: 'گوشی هوشمند سامسونگ',
    productCategory: 'الکترونیک',
    content: 'این گوشی واقعاً عالی است! کیفیت دوربین فوق‌العاده و باتری طولانی مدت دارد.',
    hasPhoto: true,
    hasVideo: false,
    isFirstReview: true,
    quality: 'excellent',
    score: 85,
    status: 'analyzed',
    createdAt: '2024-01-15T10:30:00Z'
  }
])

// Sample referrals data
const referrals = ref([
  {
    id: 1,
    referrerName: 'علی احمدی',
    referrerEmail: 'ali@example.com',
    referrerAvatar: '/avatars/ali.jpg',
    referredName: 'فاطمه محمدی',
    referredEmail: 'fateme@example.com',
    referredAvatar: '/avatars/fateme.jpg',
    referralCode: 'ALI2024',
    codeType: 'شخصی',
    createdAt: '2024-01-15T10:30:00Z',
    status: 'successful',
    quality: 'high',
    score: 70
  }
])

// Sample scenarios data (for ScenarioManagementSystem)
const scenarios = ref([
  {
    id: 1,
    name: 'سناریو پایه فروشگاه',
    description: 'سناریو استاندارد برای فروشگاه‌های آنلاین',
    type: 'basic',
    status: 'active',
    createdAt: '2024-01-15T10:30:00Z',
    affectedUsers: 1250,
    performance: 85
  }
])

// Settings data
const scoringSettings = ref({
  isEnabled: true,
  scoringRules: {
    loginScore: 1,
    activityTimeScore: 0.5,
    profileCompletionScore: 10,
    purchaseScore: 10,
    referralScore: 20,
    surveyScore: 3,
    reviewScore: 5,
    productRatingScore: 2,
    answerScore: 3,
    discussionScore: 2,
    positiveFeedbackScore: 5,
    shareScore: 2,
    inappropriateBehaviorPenalty: 10,
    inappropriateReviewPenalty: 5,
    inactivityPenalty: 1,
    violationPenalty: 50,
    dailyScoreLimit: 100,
    dailyReviewLimit: 10,
    dailyReferralLimit: 5
  }
})

const systemSettings = ref({
  isSystemEnabled: true,
  autoScoringEnabled: true,
  autoNotifications: true,
  showScoreToUsers: true,
  showLevelToUsers: true,
  warningSystemEnabled: true,
  minTopUserScore: 500,
  autoUpgradeThreshold: 100,
  blockThreshold: -50,
  warningThreshold: -20,
  dailyScoreLimit: 100,
  dailyReviewLimit: 10,
  dailyReferralLimit: 5,
  monthlyScoreLimit: 2000,
  monthlyPurchaseLimit: 50,
  monthlyActivityLimit: 500,
  maxTotalScore: 50000,
  minAccountAge: 7,
  minActivities: 10,
  notifications: {
    scoreChange: true,
    levelUpgrade: true,
    warning: true,
    blocking: true
  },
  adminNotifications: {
    newTopUser: true,
    userBlocked: true,
    systemAlert: true,
    dailyReport: false
  }
})

// Event handlers
const handleViewUserDetails = (user: any) => {
  console.log('مشاهده جزئیات کاربر:', user)
  // Navigate to user details page
  navigateTo(`/admin/users/${user.id}`)
}

const handleEditUser = (user: any) => {
  console.log('ویرایش کاربر:', user)
  // Open edit user modal
}

const handleToggleUserStatus = (user: any) => {
  console.log('تغییر وضعیت کاربر:', user)
  user.status = user.status === 'blocked' ? 'active' : 'blocked'
  // API call to update user status
}

const handleSaveScoringSettings = (settings: any) => {
  console.log('ذخیره تنظیمات امتیازدهی:', settings)
  scoringSettings.value = settings
  // API call to save settings
}

const handleSubmitManualScore = (data: any) => {
  console.log('اعمال امتیاز دستی:', data)
  // API call to apply manual score
}

const handleExportHistory = (data: any) => {
  console.log('خروجی تاریخچه:', data)
  // Export history to Excel
}

const handleSendAlert = (data: any) => {
  console.log('ارسال هشدار:', data)
  // API call to send alert
}

const handleSendBulkAlerts = (data: any) => {
  console.log('ارسال هشدار گروهی:', data)
  // API call to send bulk alerts
}

const handleExportAlertHistory = (data: any) => {
  console.log('خروجی تاریخچه هشدارها:', data)
  // Export alert history to Excel
}

const handleSaveBenefits = (data: any) => {
  console.log('ذخیره مزایا:', data)
  // API call to save benefits
}

const handleExportBenefitsData = (data: any) => {
  console.log('خروجی داده‌های مزایا:', data)
  // Export benefits data to Excel
}

const handleSaveSystemSettings = (settings: any) => {
  console.log('ذخیره تنظیمات سیستم:', settings)
  systemSettings.value = settings
  // API call to save system settings
}

const handleExportData = (type: string, data: any) => {
  console.log('خروجی داده:', type, data)
  // Export data to Excel based on type
}

const exportData = () => {
  console.log('خروجی کلی داده‌ها')
  // Export all data to Excel
}

// Review Quality System handlers
const handleSaveReviewSettings = (settings: any) => {
  console.log('ذخیره تنظیمات کیفیت نظرات:', settings)
  // API call to save review settings
}

const handleAnalyzeReviews = (reviews: any[]) => {
  console.log('تحلیل نظرات:', reviews)
  // API call to analyze reviews
}

const handleExportReviewReport = (data: any) => {
  console.log('خروجی گزارش کیفیت نظرات:', data)
  // Export review quality report
}

// Referral System handlers
const handleSaveReferralSettings = (settings: any) => {
  console.log('ذخیره تنظیمات سیستم ارجاع:', settings)
  // API call to save referral settings
}

const handleGenerateReferralCodes = (users: any[]) => {
  console.log('تولید کدهای ارجاع:', users)
  // Generate referral codes for users
}

const handleExportReferralReport = (data: any) => {
  console.log('خروجی گزارش ارجاع:', data)
  // Export referral report
}

// Purchase Continuity System handlers
const handleSaveContinuitySettings = (settings: any) => {
  console.log('ذخیره تنظیمات تداوم خرید:', settings)
  // API call to save continuity settings
}

const handleAnalyzePurchasePatterns = (users: any[]) => {
  console.log('تحلیل الگوهای خرید:', users)
  // Analyze purchase patterns
}

const handleExportContinuityReport = (data: any) => {
  console.log('خروجی گزارش تداوم خرید:', data)
  // Export continuity report
}

// Return Rate System handlers
const handleSaveReturnSettings = (settings: any) => {
  console.log('ذخیره تنظیمات نرخ بازگشت:', settings)
  // API call to save return settings
}

const handleAnalyzeReturnPatterns = (users: any[]) => {
  console.log('تحلیل الگوهای بازگشت:', users)
  // Analyze return patterns
}

const handleExportReturnReport = (data: any) => {
  console.log('خروجی گزارش نرخ بازگشت:', data)
  // Export return report
}

// Account Age System handlers
const handleSaveAgeSettings = (settings: any) => {
  console.log('ذخیره تنظیمات قدمت حساب:', settings)
  // API call to save age settings
}

const handleAnalyzeAccountAges = (users: any[]) => {
  console.log('تحلیل قدمت حساب‌ها:', users)
  // Analyze account ages
}

const handleExportAgeReport = (data: any) => {
  console.log('خروجی گزارش قدمت حساب:', data)
  // Export age report
}

// Scenario Management System handlers
const handleSaveScenario = (scenario: any) => {
  console.log('ذخیره سناریو:', scenario)
  // API call to save scenario
}

const handleUpdateScenario = (scenario: any) => {
  console.log('به‌روزرسانی سناریو:', scenario)
  // API call to update scenario
}

const handleDeleteScenario = (scenario: any) => {
  console.log('حذف سناریو:', scenario)
  // API call to delete scenario
}

const handleExportScenarios = (data: any) => {
  console.log('خروجی سناریوها:', data)
  // Export scenarios
}

// Setup and Testing System handlers
const handleRunSetup = (steps: any[]) => {
  console.log('اجرای راه‌اندازی:', steps)
  // Run setup steps
}

const handleRunTests = (tests: any[]) => {
  console.log('اجرای تست‌ها:', tests)
  // Run tests
}

const handleViewTestResults = (results: any) => {
  console.log('مشاهده نتایج تست:', results)
  // View test results
}

// Lifecycle
onMounted(() => {
  // Load initial data from API
  console.log('سیستم رتبه‌بندی کاربران بارگذاری شد')
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 