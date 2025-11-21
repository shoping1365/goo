<template>
  <div class="bg-white" dir="rtl">
    <!-- Header Section -->
    <div class="bg-white shadow-lg border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center space-x-4 space-x-reverse">
            <div class="p-3 bg-gradient-to-r from-blue-500 to-purple-600 rounded-xl shadow-lg">
              <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
              </svg>
            </div>
            <div>
              <h1 class="text-3xl font-bold text-gray-900">نظر سنجی بعد از خرید</h1>
              <p class="text-gray-600 mt-1">مدیریت و تحلیل نظرات مشتریان پس از خرید</p>
            </div>
          </div>
          <div class="flex items-center space-x-3 space-x-reverse">
            <button class="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
              </svg>
              <span>ایجاد نظرسنجی جدید</span>
            </button>
            
            <button class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors flex items-center space-x-2 space-x-reverse">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              <span>گزارش‌ها</span>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content with Sidebar -->
    <div class="flex">
      <!-- Sidebar Navigation -->
      <div class="w-64 bg-white shadow-lg border-l border-gray-200">
        <div class="p-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-4">منوی نظرسنجی</h3>
          <nav class="space-y-2">
            <button 
              v-for="(section, index) in sidebarSections" 
              :key="index"
              :class="[
                'w-full text-right px-4 py-3 rounded-lg transition-all duration-200 flex items-center justify-between',
                activeSection === section.id 
                  ? 'bg-blue-100 text-blue-700 border-r-4 border-blue-500' 
                  : 'text-gray-700 hover:bg-gray-100'
              ]"
              @click="activeSection = section.id"
            >
              <span class="text-sm font-medium">{{ section.title }}</span>
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
              </svg>
            </button>
          </nav>
        </div>
      </div>

      <!-- Main Content Area -->
      <div class="flex-1 max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        
      <!-- Order Information Header -->
        <div v-if="activeSection === 'order-info'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">اطلاعات سفارش</h2>
      <OrderInfoHeader />
          </div>
        </div>
      
      <!-- Notification Banner -->
        <div v-if="activeSection === 'notification-banner'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">اعلان‌ها</h2>
      <NotificationBanner />
          </div>
        </div>
      
      <!-- Progress Bar -->
        <div v-if="activeSection === 'progress-bar'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">نوار پیشرفت</h2>
      <ProgressBar :current-step="currentStep" :steps="surveySteps" />
          </div>
        </div>
      
        <!-- Overall Experience Rating -->
        <div v-if="activeSection === 'overall-experience'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">امتیاز کلی تجربه</h2>
        <OverallExperienceRating />
          </div>
        </div>
        
        <!-- Detailed Rating -->
        <div v-if="activeSection === 'detailed-rating'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">امتیاز جزئی</h2>
        <DetailedRating />
          </div>
        </div>
        
        <!-- Interactive Star Rating -->
        <div v-if="activeSection === 'interactive-rating'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">امتیاز تعاملی</h2>
        <InteractiveStarRating />
          </div>
        </div>
        
        <!-- Free Text Comments -->
        <div v-if="activeSection === 'free-text'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">نظرات آزاد</h2>
        <FreeTextComments />
          </div>
        </div>
        
        <!-- Specialized Questions -->
        <div v-if="activeSection === 'specialized-questions'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">سوالات تخصصی</h2>
        <SpecializedQuestions />
          </div>
        </div>
        
        <!-- Optional Section -->
        <div v-if="activeSection === 'optional-section'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">بخش اختیاری</h2>
        <OptionalSection />
          </div>
        </div>
        
        <!-- Reward Section -->
        <div v-if="activeSection === 'reward-section'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">بخش پاداش</h2>
        <RewardSection />
          </div>
        </div>
        
        <!-- Thank You Message -->
        <div v-if="activeSection === 'thank-you'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">پیام تشکر</h2>
        <ThankYouMessage />
          </div>
        </div>
        
        <!-- Submit Section -->
        <div v-if="activeSection === 'submit-section'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">بخش ارسال</h2>
        <SubmitSection />
          </div>
      </div>
      
        <!-- Survey Stats Dashboard -->
        <div v-if="activeSection === 'stats-dashboard'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">داشبورد آمار</h2>
          <SurveyStatsDashboard />
          </div>
        </div>
        
        <!-- Rating Stats -->
        <div v-if="activeSection === 'rating-stats'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">آمار امتیازات</h2>
          <RatingStats />
          </div>
        </div>
        
        <!-- Survey Reports -->
        <div v-if="activeSection === 'survey-reports'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">گزارش‌های نظرسنجی</h2>
          <SurveyReports />
          </div>
        </div>
        
        <!-- Notification System -->
        <div v-if="activeSection === 'notification-system'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">سیستم اعلان</h2>
          <NotificationSystem />
          </div>
        </div>
        
        <!-- Validation Preview -->
        <div v-if="activeSection === 'validation-preview'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">پیش‌نمایش اعتبارسنجی</h2>
          <ValidationPreview 
            :survey-data="surveyData" 
            :order-info="orderInfo"
            @edit="handleEdit"
            @save-draft="handleSaveDraft"
            @submit="handleSubmit"
          />
          </div>
        </div>
        
        <!-- File Upload Manager -->
        <div v-if="activeSection === 'file-upload'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">مدیریت فایل</h2>
          <FileUploadManager />
          </div>
        </div>
        
        <!-- Performance Optimizer -->
        <div v-if="activeSection === 'performance-optimizer'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">بهینه‌ساز عملکرد</h2>
          <PerformanceOptimizer />
          </div>
        </div>
        
        <!-- Security Manager -->
        <div v-if="activeSection === 'security-manager'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">مدیریت امنیت</h2>
          <SecurityManager />
          </div>
        </div>
        
        <!-- Backup Restore -->
        <div v-if="activeSection === 'backup-restore'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">پشتیبان‌گیری و بازیابی</h2>
          <BackupRestore />
        </div>
      </div>
      
        <!-- Management Tools -->
        <div v-if="activeSection === 'management-tools'" class="space-y-6">
          <div class="bg-white rounded-lg shadow-md border border-gray-200 p-6">
            <h2 class="text-xl font-bold text-gray-900 mb-4">ابزارهای مدیریتی</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div class="bg-white rounded-xl shadow-lg border border-gray-200 p-6 hover:shadow-xl transition-shadow">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-blue-100 rounded-lg">
              <svg class="w-6 h-6 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-2">تحلیل داده‌ها</h3>
          </div>
          <p class="text-gray-600 text-sm mb-4">تحلیل پیشرفته نظرات و گزارش‌گیری</p>
          <NuxtLink to="./DataAnalytics" class="text-blue-600 hover:text-blue-800 text-sm font-medium">مشاهده تحلیل‌ها →</NuxtLink>
        </div>
        
        <div class="bg-white rounded-xl shadow-lg border border-gray-200 p-6 hover:shadow-xl transition-shadow">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-green-100 rounded-lg">
              <svg class="w-6 h-6 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-2">مدیریت نظرات</h3>
          </div>
          <p class="text-gray-600 text-sm mb-4">بررسی و تأیید نظرات دریافتی</p>
          <NuxtLink to="./SurveyReviewManagement" class="text-green-600 hover:text-green-800 text-sm font-medium">مدیریت نظرات →</NuxtLink>
        </div>
        
        <div class="bg-white rounded-xl shadow-lg border border-gray-200 p-6 hover:shadow-xl transition-shadow">
          <div class="flex items-center mb-4">
            <div class="p-2 bg-purple-100 rounded-lg">
              <svg class="w-6 h-6 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z"></path>
              </svg>
            </div>
            <h3 class="text-lg font-semibold text-gray-900 mr-2">یکپارچه‌سازی</h3>
          </div>
          <p class="text-gray-600 text-sm mb-4">اتصال به سرویس‌های خارجی</p>
          <NuxtLink to="./IntegrationManager" class="text-purple-600 hover:text-purple-800 text-sm font-medium">تنظیمات اتصال →</NuxtLink>
              </div>
            </div>
          </div>
        </div>
        
        <!-- Default message when no section is selected -->
        <div v-if="!activeSection" class="flex items-center justify-center h-64">
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
declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
declare const useHead: (head: { title?: string; meta?: Array<{ name?: string; content?: string }> }) => void
</script>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import BackupRestore from './BackupRestore.vue'
import DetailedRating from './DetailedRating.vue'
import FileUploadManager from './FileUploadManager.vue'
import FreeTextComments from './FreeTextComments.vue'
import InteractiveStarRating from './InteractiveStarRating.vue'
import NotificationBanner from './NotificationBanner.vue'
import NotificationSystem from './NotificationSystem.vue'
import OptionalSection from './OptionalSection.vue'
import OrderInfoHeader from './OrderInfoHeader.vue'
import OverallExperienceRating from './OverallExperienceRating.vue'
import PerformanceOptimizer from './PerformanceOptimizer.vue'
import ProgressBar from './ProgressBar.vue'
import RatingStats from './RatingStats.vue'
import RewardSection from './RewardSection.vue'
import SecurityManager from './SecurityManager.vue'
import SpecializedQuestions from './SpecializedQuestions.vue'
import SubmitSection from './SubmitSection.vue'
import SurveyReports from './SurveyReports.vue'
import SurveyStatsDashboard from './SurveyStatsDashboard.vue'
import ThankYouMessage from './ThankYouMessage.vue'
import ValidationPreview from './ValidationPreview.vue'

// Reactive data
const currentStep = ref(0)
const activeSection = ref('order-info')

const sidebarSections = ref([
  { id: 'order-info', title: 'اطلاعات سفارش' },
  { id: 'notification-banner', title: 'اعلان‌ها' },
  { id: 'progress-bar', title: 'نوار پیشرفت' },
  { id: 'overall-experience', title: 'امتیاز کلی' },
  { id: 'detailed-rating', title: 'امتیاز جزئی' },
  { id: 'interactive-rating', title: 'امتیاز تعاملی' },
  { id: 'free-text', title: 'نظرات آزاد' },
  { id: 'specialized-questions', title: 'سوالات تخصصی' },
  { id: 'optional-section', title: 'بخش اختیاری' },
  { id: 'reward-section', title: 'بخش پاداش' },
  { id: 'thank-you', title: 'پیام تشکر' },
  { id: 'submit-section', title: 'بخش ارسال' },
  { id: 'stats-dashboard', title: 'داشبورد آمار' },
  { id: 'rating-stats', title: 'آمار امتیازات' },
  { id: 'survey-reports', title: 'گزارش‌های نظرسنجی' },
  { id: 'notification-system', title: 'سیستم اعلان' },
  { id: 'validation-preview', title: 'پیش‌نمایش اعتبارسنجی' },
  { id: 'file-upload', title: 'مدیریت فایل' },
  { id: 'performance-optimizer', title: 'بهینه‌ساز عملکرد' },
  { id: 'security-manager', title: 'مدیریت امنیت' },
  { id: 'backup-restore', title: 'پشتیبان‌گیری' },
  { id: 'management-tools', title: 'ابزارهای مدیریتی' }
])

const surveySteps = ref([
  { title: 'اطلاعات سفارش', description: 'بررسی جزئیات سفارش', completed: false },
  { title: 'امتیاز کلی', description: 'امتیازدهی کلی تجربه خرید', completed: false },
  { title: 'امتیاز جزئی', description: 'امتیازدهی جنبه‌های مختلف', completed: false },
  { title: 'نظرات', description: 'نظرات و پیشنهادات', completed: false },
  { title: 'سوالات تخصصی', description: 'سوالات تخصصی', completed: false },
  { title: 'بخش اختیاری', description: 'اطلاعات اضافی', completed: false },
  { title: 'پاداش', description: 'دریافت پاداش', completed: false },
  { title: 'ارسال', description: 'ارسال نهایی', completed: false }
])

const surveyData = reactive({
  overallRating: 0,
  detailedRatings: {
    productQuality: 0,
    pricing: 0,
    deliverySpeed: 0,
    packaging: 0,
    afterSales: 0,
    userInterface: 0
  },
  comments: {
    overallExperience: '',
    strengths: '',
    weaknesses: '',
    improvements: ''
  },
  specializedQuestions: {
    paymentSatisfaction: '',
    shippingSatisfaction: '',
    communicationQuality: '',
    productInfoAccuracy: '',
    supportResponseSpeed: ''
  },
  optional: {
    newsletter: false,
    contactInfo: ''
  }
})

const orderInfo = reactive({
  orderNumber: 'ORD-2024-001',
  orderDate: '2024-01-15'
})

// Event handlers
const handleEdit = () => {
}

const handleSaveDraft = (data: Record<string, unknown>) => {
}

const handleSubmit = () => {
}

/* const scrollToSection = (sectionId: string) => {
  activeSection.value = sectionId
  const element = document.getElementById(sectionId)
  if (element) {
    element.scrollIntoView({ behavior: 'smooth', block: 'start' })
  }
} */

// Page metadata
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

useHead({
  title: 'نظر سنجی بعد از خرید - پنل مدیریت',
  meta: [
    { name: 'description', content: 'مدیریت و تحلیل نظرات مشتریان پس از خرید' }
  ]
})
</script>

<style scoped>
/* Custom scrollbar */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a1a1a1;
}

/* Smooth animations */
.space-y-8 > * {
  animation: fadeInUp 0.6s ease-out;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Stagger animation for grid items */
.grid > * {
  animation: fadeInUp 0.6s ease-out;
}

.grid > *:nth-child(1) { animation-delay: 0s; }
.grid > *:nth-child(2) { animation-delay: 0s; }
.grid > *:nth-child(3) { animation-delay: 0s; }
</style> 
