<template>
  <div class="flex-1 p-6 space-y-8">
    <!-- تنظیمات ترب -->
    <TorbSettings />
    
    <!-- تنظیمات ایمالز -->
    <EmailsSettings />
    
    <!-- گزارشات ترب -->
    <TorbReports />
    
    <!-- گزارشات ایمالز -->
    <EmailsReports />
    
    <!-- گزارش بودجه ترب -->
    <TorbBudgetReport />
    
    <!-- گزارش بودجه ایمالز -->
    <EmailsBudgetReport />
    
    <!-- کرالر ترب -->
    <TorbCrawler />
    
    <!-- کرالر ایمالز -->
    <EmailsCrawler />
    
    <!-- مدیریت محتوا ترب -->
    <TorbContentManagement />
    
    <!-- مدیریت محتوا ایمالز -->
    <EmailsContentManagement />
    
    <!-- تحلیل‌های پیشرفته ترب -->
    <TorbAdvancedAnalytics />
    
    <!-- تحلیل‌های پیشرفته ایمالز -->
    <EmailsAdvancedAnalytics />
    
    <!-- مدیریت اعلان‌ها ترب -->
    <TorbNotificationManagement />
    
    <!-- مدیریت اعلان‌ها ایمالز -->
    <EmailsNotificationManagement />
    
    <!-- یکپارچه‌سازی سیستم ترب -->
    <TorbSystemIntegration />
    
    <!-- یکپارچه‌سازی سیستم ایمالز -->
    <EmailsSystemIntegration />
    
    <!-- مدیریت API ترب -->
    <TorbApiManagement />
    
    <!-- مدیریت API ایمالز -->
    <EmailsApiManagement />
    
    <!-- تنظیمات پیشرفته ترب -->
    <TorbAdvancedSettings />
    
    <!-- تنظیمات پیشرفته ایمالز -->
    <EmailsAdvancedSettings />
    
    <!-- گزارش‌های جامع ترب -->
    <TorbComprehensiveReports />
    
    <!-- گزارش‌های جامع ایمالز -->
    <EmailsComprehensiveReports />
    
    <!-- مدیریت رقبا ترب -->
    <TorbCompetitorManagement />
    
    <!-- مدیریت رقبا ایمالز -->
    <EmailsCompetitorManagement />
    
    <!-- نظرسنجی کاربران ترب -->
    <TorbUserSurvey />
    
    <!-- نظرسنجی کاربران ایمالز -->
    <EmailsUserSurvey />
    
    <!-- مدیریت کمپین‌ها ترب -->
    <TorbCampaignManagement />
    
    <!-- مدیریت کمپین‌ها ایمالز -->
    <EmailsCampaignManagement />
    
    <!-- مدیریت محصولات ترب -->
    <TorbProductManagement />
    
    <!-- مدیریت محصولات ایمالز -->
    <EmailsProductManagement />
    
    <!-- مدیریت قیمت‌ها ترب -->
    <TorbPriceManagement />
    
    <!-- مدیریت قیمت‌ها ایمالز -->
    <EmailsPriceManagement />
    
    <!-- مدیریت کاربران ترب -->
    <TorbUserManagement />
    
    <!-- مدیریت کاربران ایمالز -->
    <EmailsUserManagement />
    
    <!-- مدیریت ترافیک ترب -->
    <TorbTrafficManagement />
    
    <!-- مدیریت ترافیک ایمالز -->
    <EmailsTrafficManagement />
    
    <!-- مدیریت SEO ترب -->
    <TorbSeoManagement />
    
    <!-- مدیریت SEO ایمالز -->
    <EmailsSeoManagement />
    
    <!-- مدیریت تبلیغات ترب -->
    <TorbAdvertisingManagement />
    
    <!-- مدیریت تبلیغات ایمالز -->
    <EmailsAdvertisingManagement />
    
    <!-- مدیریت بازاریابی ترب -->
    <TorbMarketingManagement />
    
    <!-- مدیریت بازاریابی ایمالز -->
    <EmailsMarketingManagement />
    
    <!-- مدیریت فروش ترب -->
    <TorbSalesManagement />
    
    <!-- مدیریت فروش ایمالز -->
    <EmailsSalesManagement />
    
    <!-- مدیریت پشتیبانی ترب -->
    <TorbSupportManagement />
    
    <!-- مدیریت پشتیبانی ایمالز -->
    <EmailsSupportManagement />
    
    <!-- مدیریت کیفیت ترب -->
    <TorbQualityManagement />
    
    <!-- مدیریت کیفیت ایمالز -->
    <EmailsQualityManagement />
    
    <!-- مدیریت امنیت ترب -->
    <TorbSecurityManagement />
    
    <!-- مدیریت امنیت ایمالز -->
    <EmailsSecurityManagement />
    
    <!-- مدیریت پشتیبان‌گیری ترب -->
    <TorbBackupManagement />
    
    <!-- مدیریت پشتیبان‌گیری ایمالز -->
    <EmailsBackupManagement />
    
    <!-- مدیریت بهینه‌سازی ترب -->
    <TorbOptimizationManagement />
    
    <!-- مدیریت بهینه‌سازی ایمالز -->
    <EmailsOptimizationManagement />
    
    <!-- مدیریت مانیتورینگ ترب -->
    <TorbMonitoringManagement />
    
    <!-- مدیریت مانیتورینگ ایمالز -->
    <EmailsMonitoringManagement />
    
    <!-- مدیریت آمار ترب -->
    <TorbStatisticsManagement />
    
    <!-- مدیریت آمار ایمالز -->
    <EmailsStatisticsManagement />
    
    <!-- مدیریت گزارش‌گیری ترب -->
    <TorbReportingManagement />
    
    <!-- مدیریت گزارش‌گیری ایمالز -->
    <EmailsReportingManagement />
    
    <!-- مدیریت داشبورد ترب -->
    <TorbDashboardManagement />
    
    <!-- مدیریت داشبورد ایمالز -->
    <EmailsDashboardManagement />
    
    <!-- مدیریت تنظیمات عمومی ترب -->
    <TorbGeneralSettings />
    
    <!-- مدیریت تنظیمات عمومی ایمالز -->
    <EmailsGeneralSettings />
  </div>
</template>

<script lang="ts">
declare const definePageMeta: (meta: { layout?: string; middleware?: string | string[] }) => void
</script>

<script setup lang="ts">
// ایمپورت تمام کامپوننت‌های ترب
import TorbAdvancedAnalytics from './components/TorbAdvancedAnalytics.vue'
import TorbAdvancedSettings from './components/TorbAdvancedSettings.vue'
import TorbAdvertisingManagement from './components/TorbAdvertisingManagement.vue'
import TorbApiManagement from './components/TorbApiManagement.vue'
import TorbBackupManagement from './components/TorbBackupManagement.vue'
import TorbBudgetReport from './components/TorbBudgetReport.vue'
import TorbCampaignManagement from './components/TorbCampaignManagement.vue'
import TorbCompetitorManagement from './components/TorbCompetitorManagement.vue'
import TorbComprehensiveReports from './components/TorbComprehensiveReports.vue'
import TorbContentManagement from './components/TorbContentManagement.vue'
import TorbCrawler from './components/TorbCrawler.vue'
import TorbDashboardManagement from './components/TorbDashboardManagement.vue'
import TorbGeneralSettings from './components/TorbGeneralSettings.vue'
import TorbMarketingManagement from './components/TorbMarketingManagement.vue'
import TorbMonitoringManagement from './components/TorbMonitoringManagement.vue'
import TorbNotificationManagement from './components/TorbNotificationManagement.vue'
import TorbOptimizationManagement from './components/TorbOptimizationManagement.vue'
import TorbPriceManagement from './components/TorbPriceManagement.vue'
import TorbProductManagement from './components/TorbProductManagement.vue'
import TorbQualityManagement from './components/TorbQualityManagement.vue'
import TorbReportingManagement from './components/TorbReportingManagement.vue'
import TorbReports from './components/TorbReports.vue'
import TorbSalesManagement from './components/TorbSalesManagement.vue'
import TorbSecurityManagement from './components/TorbSecurityManagement.vue'
import TorbSeoManagement from './components/TorbSeoManagement.vue'
import TorbSettings from './components/TorbSettings.vue'
import TorbStatisticsManagement from './components/TorbStatisticsManagement.vue'
import TorbSupportManagement from './components/TorbSupportManagement.vue'
import TorbSystemIntegration from './components/TorbSystemIntegration.vue'
import TorbTrafficManagement from './components/TorbTrafficManagement.vue'
import TorbUserManagement from './components/TorbUserManagement.vue'
import TorbUserSurvey from './components/TorbUserSurvey.vue'

// ایمپورت تمام کامپوننت‌های ایمالز
import EmailsAdvancedAnalytics from './components/EmailsAdvancedAnalytics.vue'
import EmailsAdvancedSettings from './components/EmailsAdvancedSettings.vue'
import EmailsAdvertisingManagement from './components/EmailsAdvertisingManagement.vue'
import EmailsApiManagement from './components/EmailsApiManagement.vue'
import EmailsBackupManagement from './components/EmailsBackupManagement.vue'
import EmailsBudgetReport from './components/EmailsBudgetReport.vue'
import EmailsCampaignManagement from './components/EmailsCampaignManagement.vue'
import EmailsCompetitorManagement from './components/EmailsCompetitorManagement.vue'
import EmailsComprehensiveReports from './components/EmailsComprehensiveReports.vue'
import EmailsContentManagement from './components/EmailsContentManagement.vue'
import EmailsCrawler from './components/EmailsCrawler.vue'
import EmailsDashboardManagement from './components/EmailsDashboardManagement.vue'
import EmailsGeneralSettings from './components/EmailsGeneralSettings.vue'
import EmailsMarketingManagement from './components/EmailsMarketingManagement.vue'
import EmailsMonitoringManagement from './components/EmailsMonitoringManagement.vue'
import EmailsNotificationManagement from './components/EmailsNotificationManagement.vue'
import EmailsOptimizationManagement from './components/EmailsOptimizationManagement.vue'
import EmailsPriceManagement from './components/EmailsPriceManagement.vue'
import EmailsProductManagement from './components/EmailsProductManagement.vue'
import EmailsQualityManagement from './components/EmailsQualityManagement.vue'
import EmailsReportingManagement from './components/EmailsReportingManagement.vue'
import EmailsReports from './components/EmailsReports.vue'
import EmailsSalesManagement from './components/EmailsSalesManagement.vue'
import EmailsSecurityManagement from './components/EmailsSecurityManagement.vue'
import EmailsSeoManagement from './components/EmailsSeoManagement.vue'
import EmailsSettings from './components/EmailsSettings.vue'
import EmailsStatisticsManagement from './components/EmailsStatisticsManagement.vue'
import EmailsSupportManagement from './components/EmailsSupportManagement.vue'
import EmailsSystemIntegration from './components/EmailsSystemIntegration.vue'
import EmailsTrafficManagement from './components/EmailsTrafficManagement.vue'
import EmailsUserManagement from './components/EmailsUserManagement.vue'
import EmailsUserSurvey from './components/EmailsUserSurvey.vue'

definePageMeta({ layout: 'admin-main', middleware: 'admin' })
</script> 
