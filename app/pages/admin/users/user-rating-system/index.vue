<template>
  <div class="min-h-screen">
    <!-- Header Section -->
    <div class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4 space-x-reverse">
            <button 
              class="p-2 rounded-md text-gray-500 hover:text-gray-700 hover:bg-gray-100 transition-colors"
              @click="toggleSidebar"
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
            <button class="bg-blue-600 text-white px-3 py-2 rounded-md hover:bg-blue-700 transition-colors text-sm" @click="activeTab = 'settings'">
              تنظیمات سیستم
            </button>
            <button class="bg-green-600 text-white px-3 py-2 rounded-md hover:bg-green-700 transition-colors text-sm" @click="exportData">
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
        class="fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden"
        @click="toggleSidebar"
      ></div>
      
      <!-- Sidebar -->
      <div 
        :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full'"
        class="fixed inset-y-0 right-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out lg:relative lg:translate-x-0"
      >
        <div class="flex items-center justify-between p-4 border-b border-gray-200">
          <h2 class="text-lg font-semibold text-gray-900">منوها</h2>
          <button 
            class="lg:hidden p-1 rounded-md text-gray-500 hover:text-gray-700 hover:bg-gray-100"
            @click="toggleSidebar"
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
              :class="activeTab === 'dashboard' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'dashboard'"
            >
              داشبورد
            </button>
            <button 
              :class="activeTab === 'users' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'users'"
            >
              مدیریت کاربران
            </button>
            <button 
              :class="activeTab === 'scoring' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'scoring'"
            >
              سیستم امتیازدهی
            </button>
            <button 
              :class="activeTab === 'manual' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'manual'"
            >
              امتیازدهی دستی
            </button>
            <button 
              :class="activeTab === 'alerts' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'alerts'"
            >
              سیستم هشدار
            </button>
            <button 
              :class="activeTab === 'benefits' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'benefits'"
            >
              مزایای کاربران برتر
            </button>
          </div>

          <!-- Advanced Menu Items -->
          <div class="space-y-1 pt-4 border-t border-gray-200">
            <h3 class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">منوهای پیشرفته</h3>
            <button 
              :class="activeTab === 'reviews' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'reviews'"
            >
              کیفیت نظرات
            </button>
            <button 
              :class="activeTab === 'referrals' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'referrals'"
            >
              سیستم ارجاع
            </button>
            <button 
              :class="activeTab === 'purchases' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'purchases'"
            >
              تداوم خرید
            </button>
            <button 
              :class="activeTab === 'returns' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'returns'"
            >
              نرخ بازگشت
            </button>
            <button 
              :class="activeTab === 'account-age' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'account-age'"
            >
              قدمت حساب
            </button>
          </div>

          <!-- System Menu Items -->
          <div class="space-y-1 pt-4 border-t border-gray-200">
            <h3 class="text-xs font-semibold text-gray-500 uppercase tracking-wider mb-2">سیستم</h3>
            <button 
              :class="activeTab === 'scenarios' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'scenarios'"
            >
              سناریوها
            </button>
            <button 
              :class="activeTab === 'setup-test' ? 'bg-blue-50 text-blue-700 border-blue-200' : 'text-gray-700 hover:bg-gray-50 border-transparent'"
              class="w-full text-right px-3 py-2 rounded-md border text-sm font-medium transition-colors"
              @click="activeTab = 'setup-test'"
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
        <ReportsAndAnalytics :users="users as unknown as Array<{ id: number; name: string; email: string; avatar: string; score: number; monthlyReviews: number; monthlyPurchases: number; monthlyReferrals: number; lastActivity: string; [key: string]: unknown }>" @export-data="handleExportData" />
      </div>

      <!-- Users Management Tab -->
      <div v-if="activeTab === 'users'">
        <UserManagement 
          :users="users as unknown as Array<{ id: number; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string }>" 
          @view-details="handleViewUserDetails as unknown as (user: { id: number; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string }) => void"
          @edit-user="handleEditUser as unknown as (user: { id: number; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string }) => void"
          @toggle-status="handleToggleUserStatus as unknown as (user: { id: number; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string }) => void"
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
          :users="users as any"
          @submit-score="handleSubmitManualScore as any"
          @export-history="handleExportHistory as any"
        />
      </div>

      <!-- Alert System Tab -->
      <div v-if="activeTab === 'alerts'">
        <AlertSystem 
          :users="users as any"
          @send-alert="handleSendAlert as any"
          @send-bulk-alerts="handleSendBulkAlerts as any"
          @export-history="handleExportAlertHistory as any"
        />
      </div>

      <!-- Top User Benefits Tab -->
      <div v-if="activeTab === 'benefits'">
        <TopUserBenefits
          :users="users as any"
          @save-benefits="handleSaveBenefits as any"
          @export-data="handleExportBenefitsData as any"
        />
      </div>

      <!-- Review Quality System Tab -->
      <div v-if="activeTab === 'reviews'">
        <ReviewQualitySystem
          :reviews="reviews"
          @save-settings="handleSaveReviewSettings as (settings: AISettings) => void"
          @analyze-reviews="handleAnalyzeReviews as (reviews: Review[]) => void"
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
          :users="users as any"
          @save-settings="handleSaveContinuitySettings as any"
          @analyze-patterns="handleAnalyzePurchasePatterns as any"
          @export-report="handleExportContinuityReport as any"
        />
      </div>

      <!-- Return Rate System Tab -->
      <div v-if="activeTab === 'returns'">
        <ReturnRateSystem
          :users="users.map(u => ({
            id: Number(u.id),
            userName: u.name || u.userName || '',
            userEmail: u.email || u.userEmail || '',
            userAvatar: u.avatar || u.userAvatar || '',
            totalPurchases: u.monthlyPurchases || 0,
            totalReturns: 0,
            returnRate: u.returnRate || 0,
            lastReturn: '',
            status: u.status || 'active',
            score: u.score,
            name: u.name,
            email: u.email,
            avatar: u.avatar
          } as ReturnPattern))"
          @save-settings="handleSaveReturnSettings"
          @analyze-patterns="handleAnalyzeReturnPatterns"
          @export-report="handleExportReturnReport"
        />
      </div>

      <!-- Account Age System Tab -->
      <div v-if="activeTab === 'account-age'">
        <AccountAgeSystem
          :users="users as any"
          @save-settings="handleSaveAgeSettings as any"
          @analyze-ages="handleAnalyzeAccountAges as any"
          @export-report="handleExportAgeReport as any"
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
          @run-setup="handleRunSetup as (steps: Step[]) => void"
          @run-tests="handleRunTests as (tests: TestEnvironment[]) => void"
          @view-results="handleViewTestResults as (results: TestResult) => void"
        />
      </div>
    </div>
  </div>
</div>
</div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
declare const navigateTo: (path: string) => Promise<void>
</script>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import AccountAgeSystem from './components/AccountAgeSystem.vue'
import AlertSystem from './components/AlertSystem.vue'
import AutoScoringSystem from './components/AutoScoringSystem.vue'
import ManualScoring from './components/ManualScoring.vue'
import PurchaseContinuitySystem from './components/PurchaseContinuitySystem.vue'
import ReferralSystem from './components/ReferralSystem.vue'
import ReportsAndAnalytics from './components/ReportsAndAnalytics.vue'
import ReturnRateSystem from './components/ReturnRateSystem.vue'
import ReviewQualitySystem from './components/ReviewQualitySystem.vue'
import ScenarioManagementSystem from './components/ScenarioManagementSystem.vue'
import SetupAndTestingSystem from './components/SetupAndTestingSystem.vue'
import SystemSettings from './components/SystemSettings.vue'
import TopUserBenefits from './components/TopUserBenefits.vue'
import UserManagement from './components/UserManagement.vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })

// جامع‌ترین interface برای User که شامل تمام ویژگی‌های مورد نیاز کامپوننت‌های مختلف است
interface User {
  id: number;
  // ویژگی‌های پایه
  name?: string;
  email?: string;
  avatar?: string;
  // ویژگی‌های جایگزین (برای سازگاری با کامپوننت‌های مختلف)
  userName?: string;
  userEmail?: string;
  userAvatar?: string;
  // ویژگی‌های امتیاز
  score: number;
  level: string;
  currentScore?: number;
  // ویژگی‌های تاریخ
  lastActivity?: string;
  joinDate?: string;
  lastPurchase?: string;
  // ویژگی‌های اضافی
  status?: string;
  monthlyReviews?: number;
  monthlyPurchases?: number;
  monthlyReferrals?: number;
  riskLevel?: string;
  activeBenefits?: string[];
  lastBenefitUsage?: string;
  pattern?: string;
  purchaseInterval?: string;
  streakLevel?: number;
  age?: number;
  loyaltyLevel?: string;
  returnRate?: number;
  [key: string]: unknown;
}

// Helper interface for ReturnRateSystem
interface ReturnPattern {
  id: number;
  userName: string;
  userEmail: string;
  userAvatar: string;
  totalPurchases: number;
  totalReturns: number;
  returnRate: number;
  lastReturn: string;
  status: string;
  score: number;
  name?: string;
  email?: string;
  avatar?: string;
  [key: string]: unknown;
}

interface AISettings {
  [key: string]: unknown;
}

interface Review {
  [key: string]: unknown;
}

interface Step {
  [key: string]: unknown;
}

interface TestEnvironment {
  [key: string]: unknown;
}

interface TestResult {
  [key: string]: unknown;
}

// Reactive data
const activeTab = ref('dashboard')
const isSidebarOpen = ref(false)

// Toggle sidebar function
const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

// Sample users data with updated structure - شامل تمام ویژگی‌های مورد نیاز کامپوننت‌های مختلف
const users = ref<User[]>([
  {
    id: 1,
    name: 'علی احمدی',
    email: 'ali@example.com',
    avatar: '/avatars/ali.jpg',
    userName: 'علی احمدی',
    userEmail: 'ali@example.com',
    userAvatar: '/avatars/ali.jpg',
    score: 850,
    level: 'diamond',
    lastActivity: '2024-01-15T10:30:00Z',
    joinDate: '2023-01-15T10:30:00Z',
    status: 'active',
    currentScore: 850,
    monthlyReviews: 15,
    monthlyPurchases: 8,
    monthlyReferrals: 5,
    riskLevel: 'low',
    activeBenefits: ['discount_10', 'free_shipping'],
    lastBenefitUsage: '2024-01-14T10:30:00Z',
    lastPurchase: '2024-01-14T10:30:00Z',
    pattern: 'regular',
    purchaseInterval: 'monthly',
    streakLevel: 3,
    age: 25,
    loyaltyLevel: 'high',
    returnRate: 0.05
  },
  {
    id: 2,
    name: 'فاطمه محمدی',
    email: 'fateme@example.com',
    avatar: '/avatars/fateme.jpg',
    userName: 'فاطمه محمدی',
    userEmail: 'fateme@example.com',
    userAvatar: '/avatars/fateme.jpg',
    score: 720,
    level: 'platinum',
    lastActivity: '2024-01-14T15:45:00Z',
    joinDate: '2023-03-14T15:45:00Z',
    status: 'active',
    currentScore: 720,
    monthlyReviews: 12,
    monthlyPurchases: 6,
    monthlyReferrals: 3,
    riskLevel: 'low',
    activeBenefits: ['discount_5'],
    lastBenefitUsage: '2024-01-13T15:45:00Z',
    lastPurchase: '2024-01-13T15:45:00Z',
    pattern: 'regular',
    purchaseInterval: 'bi-weekly',
    streakLevel: 2,
    age: 28,
    loyaltyLevel: 'medium',
    returnRate: 0.08
  },
  {
    id: 3,
    name: 'محمد رضایی',
    email: 'mohammad@example.com',
    avatar: '/avatars/mohammad.jpg',
    userName: 'محمد رضایی',
    userEmail: 'mohammad@example.com',
    userAvatar: '/avatars/mohammad.jpg',
    score: -75,
    level: 'bronze',
    lastActivity: '2024-01-13T09:20:00Z',
    joinDate: '2023-05-13T09:20:00Z',
    status: 'blocked',
    currentScore: -75,
    monthlyReviews: 2,
    monthlyPurchases: 1,
    monthlyReferrals: 0,
    riskLevel: 'high',
    activeBenefits: [],
    lastPurchase: '2024-01-10T09:20:00Z',
    pattern: 'irregular',
    purchaseInterval: 'irregular',
    streakLevel: 0,
    age: 35,
    loyaltyLevel: 'low',
    returnRate: 0.25
  },
  {
    id: 4,
    name: 'زهرا کریمی',
    email: 'zahra@example.com',
    avatar: '/avatars/zahra.jpg',
    userName: 'زهرا کریمی',
    userEmail: 'zahra@example.com',
    userAvatar: '/avatars/zahra.jpg',
    score: 650,
    level: 'gold',
    lastActivity: '2024-01-15T14:10:00Z',
    joinDate: '2023-07-15T14:10:00Z',
    status: 'active',
    currentScore: 650,
    monthlyReviews: 10,
    monthlyPurchases: 5,
    monthlyReferrals: 2,
    riskLevel: 'low',
    activeBenefits: ['discount_5', 'free_shipping'],
    lastBenefitUsage: '2024-01-14T14:10:00Z',
    lastPurchase: '2024-01-14T14:10:00Z',
    pattern: 'regular',
    purchaseInterval: 'monthly',
    streakLevel: 2,
    age: 24,
    loyaltyLevel: 'medium',
    returnRate: 0.10
  },
  {
    id: 5,
    name: 'حسین نوری',
    email: 'hossein@example.com',
    avatar: '/avatars/hossein.jpg',
    userName: 'حسین نوری',
    userEmail: 'hossein@example.com',
    userAvatar: '/avatars/hossein.jpg',
    score: 150,
    level: 'silver',
    lastActivity: '2024-01-12T11:30:00Z',
    joinDate: '2023-08-12T11:30:00Z',
    status: 'active',
    currentScore: 150,
    monthlyReviews: 5,
    monthlyPurchases: 3,
    monthlyReferrals: 1,
    riskLevel: 'medium',
    activeBenefits: ['discount_5'],
    lastPurchase: '2024-01-11T11:30:00Z',
    pattern: 'regular',
    purchaseInterval: 'monthly',
    streakLevel: 1,
    age: 22,
    loyaltyLevel: 'medium',
    returnRate: 0.12
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
const handleViewUserDetails = (user: User | { id: number | string; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string; [key: string]: unknown }) => {
  // Navigate to user details page
  navigateTo(`/admin/users/${Number(user.id)}`)
}

const handleEditUser = (_user: User | { id: number | string; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string; [key: string]: unknown }) => {
  // Open edit user modal
}

const handleToggleUserStatus = (user: User | { id: number | string; name: string; email: string; avatar?: string; score: number; level: string; status: string; lastActivity: string; [key: string]: unknown }) => {
  if ('status' in user && user.status !== undefined) {
    (user as { status: string }).status = user.status === 'blocked' ? 'active' : 'blocked'
  }
  // API call to update user status
}

const handleSaveScoringSettings = (settings: Record<string, unknown>) => {
  scoringSettings.value = settings as typeof scoringSettings.value
  // API call to save settings
}

const handleSubmitManualScore = (_data: Record<string, unknown>) => {
  // API call to apply manual score
}

const handleExportHistory = (_data: Record<string, unknown>) => {
  // Export history to Excel
}

const handleSendAlert = (_data: Record<string, unknown>) => {
  // API call to send alert
}

const handleSendBulkAlerts = (_data: Record<string, unknown>) => {
  // API call to send bulk alerts
}

const handleExportAlertHistory = (_data: Record<string, unknown>) => {
  // Export alert history to Excel
}

const handleSaveBenefits = (_data: Record<string, unknown>) => {
  // API call to save benefits
}

const handleExportBenefitsData = (_type: string, _data: Record<string, unknown>) => {
  // Export benefits data to Excel
  // TODO: Implement export functionality
}

const handleSaveSystemSettings = (settings: Record<string, unknown>) => {
  systemSettings.value = settings as typeof systemSettings.value
  // API call to save system settings
}

const handleExportData = (_type: string, _data: Record<string, unknown>) => {
  // Export data to Excel based on type
}

const exportData = () => {
  // Export all data to Excel
}

// Review Quality System handlers
const handleSaveReviewSettings = (_settings: Record<string, unknown>) => {
  // API call to save review settings
}

const handleAnalyzeReviews = (_reviews: Record<string, unknown>[]) => {
  // API call to analyze reviews
}

const handleExportReviewReport = (_data: Record<string, unknown>) => {
  // Export review quality report
}

// Referral System handlers
const handleSaveReferralSettings = (_settings: Record<string, unknown>) => {
  // API call to save referral settings
}

const handleGenerateReferralCodes = (_users: Record<string, unknown>[]) => {
  // Generate referral codes for users
}

const handleExportReferralReport = (_data: Record<string, unknown>) => {
  // Export referral report
}

// Purchase Continuity System handlers
const handleSaveContinuitySettings = (_settings: Record<string, unknown>) => {
  // API call to save continuity settings
}

const handleAnalyzePurchasePatterns = (_users: Record<string, unknown>[]) => {
  // Analyze purchase patterns
}

const handleExportContinuityReport = (_data: Record<string, unknown>) => {
  // Export continuity report
}

// Return Rate System handlers
const handleSaveReturnSettings = (_settings: Record<string, unknown>) => {
  // API call to save return settings
}

const handleAnalyzeReturnPatterns = (_users: Record<string, unknown>[]) => {
  // Analyze return patterns
}

const handleExportReturnReport = (_data: Record<string, unknown>) => {
  // Export return report
}

// Account Age System handlers
const handleSaveAgeSettings = (_settings: Record<string, unknown>) => {
  // API call to save age settings
}

const handleAnalyzeAccountAges = (_users: Record<string, unknown>[]) => {
  // Analyze account ages
}

const handleExportAgeReport = (_data: Record<string, unknown>) => {
  // Export age report
}

// Scenario Management System handlers
const handleSaveScenario = (_scenario: Record<string, unknown>) => {
  // API call to save scenario
}

const handleUpdateScenario = (_scenario: Record<string, unknown>) => {
  // API call to update scenario
}

const handleDeleteScenario = (_scenario: Record<string, unknown>) => {
  // API call to delete scenario
}

const handleExportScenarios = (_data: Record<string, unknown>) => {
  // Export scenarios
}

// Setup and Testing System handlers
const handleRunSetup = (_steps: Record<string, unknown>[]) => {
  // Run setup steps
}

const handleRunTests = (_tests: Record<string, unknown>[]) => {
  // Run tests
}

const handleViewTestResults = (_results: Record<string, unknown>) => {
  // View test results
}

// Lifecycle
onMounted(() => {
  // Load initial data from API
})
</script>

<style scoped>
/* استایل‌های خاص کامپوننت */
</style> 