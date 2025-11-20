<template>
     <div class="notification-system bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
       <div class="flex items-center mb-6">
         <div class="p-3 bg-gradient-to-r from-indigo-400 to-indigo-600 rounded-xl shadow-lg">
           <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6v-2H4v2zM20 4v6h-2V4h2zM4 4v6h2V4H4z"></path>
           </svg>
         </div>
         <div class="mr-4">
           <h2 class="text-2xl font-bold text-gray-900">سیستم اطلاع‌رسانی</h2>
           <p class="text-gray-600 mt-1">مدیریت ایمیل، پیامک و اعلان‌های سایت</p>
         </div>
       </div>
   
       <!-- Notification Stats -->
       <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
         <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-blue-600">ایمیل ارسال شده</p>
               <p class="text-3xl font-bold text-blue-900">{{ notificationStats.emailsSent }}</p>
             </div>
             <div class="p-3 bg-blue-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-blue-700">
             {{ notificationStats.emailDeliveryRate }}% نرخ تحویل
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-green-600">پیامک ارسال شده</p>
               <p class="text-3xl font-bold text-green-900">{{ notificationStats.smsSent }}</p>
             </div>
             <div class="p-3 bg-green-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-green-700">
             {{ notificationStats.smsDeliveryRate }}% نرخ تحویل
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-purple-600">اعلان‌های سایت</p>
               <p class="text-3xl font-bold text-purple-900">{{ notificationStats.siteNotifications }}</p>
             </div>
             <div class="p-3 bg-purple-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-5 5v-5zM4 19h6v-2H4v2zM20 4v6h-2V4h2zM4 4v6h2V4H4z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-purple-700">
             {{ notificationStats.siteNotificationRate }}% نرخ بازدید
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-orange-50 to-orange-100 rounded-xl p-6 border border-orange-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-orange-600">نرخ پاسخ</p>
               <p class="text-3xl font-bold text-orange-900">{{ notificationStats.responseRate }}%</p>
             </div>
             <div class="p-3 bg-orange-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-orange-700">
             میانگین کلی
           </div>
         </div>
       </div>
   
       <!-- Email Settings -->
       <div class="bg-blue-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات ایمیل</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">ارسال خودکار</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input v-model="emailSettings.autoSend" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">فعال کردن ارسال خودکار</span>
               </label>
               <select v-model="emailSettings.timing" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
                 <option value="immediate">بلافاصله پس از تحویل</option>
                 <option value="1_hour">۱ ساعت بعد</option>
                 <option value="24_hours">۲۴ ساعت بعد</option>
                 <option value="3_days">۳ روز بعد</option>
               </select>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">قالب ایمیل</label>
             <div class="space-y-2">
               <select v-model="emailSettings.template" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
                 <option value="default">قالب پیش‌فرض</option>
                 <option value="custom">قالب سفارشی</option>
                 <option value="branded">قالب برندینگ</option>
               </select>
               <label class="flex items-center">
                 <input v-model="emailSettings.includeReward" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">شامل پاداش</span>
               </label>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">محتوا</label>
             <div class="space-y-2">
               <input v-model="emailSettings.subject" type="text" placeholder="موضوع ایمیل" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
               <textarea v-model="emailSettings.body" rows="4" placeholder="متن ایمیل" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"></textarea>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input v-model="emailSettings.trackOpens" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">پیگیری باز شدن</span>
               </label>
               <label class="flex items-center">
                 <input v-model="emailSettings.trackClicks" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">پیگیری کلیک</span>
               </label>
               <label class="flex items-center">
                 <input v-model="emailSettings.retryFailed" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">تلاش مجدد برای ناموفق</span>
               </label>
             </div>
           </div>
         </div>
       </div>
   
       <!-- SMS Settings -->
       <div class="bg-green-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات پیامک</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">ارسال پیامک</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input v-model="smsSettings.enabled" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">فعال کردن پیامک</span>
               </label>
               <select v-model="smsSettings.provider" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
                 <option value="kavenegar">کاوه‌نگار</option>
                 <option value="melipayamak">ملی پیامک</option>
                 <option value="smsir">اس‌ام‌اس آی‌آر</option>
                 <option value="custom">سفارشی</option>
               </select>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">متن پیامک</label>
             <div class="space-y-2">
               <textarea v-model="smsSettings.message" rows="4" placeholder="متن پیامک" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"></textarea>
               <div class="text-xs text-gray-500">
                 {{ smsSettings.message.length }}/160 کاراکتر
               </div>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">زمان ارسال</label>
             <div class="space-y-2">
               <select v-model="smsSettings.timing" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
                 <option value="immediate">بلافاصله</option>
                 <option value="1_hour">۱ ساعت بعد</option>
                 <option value="24_hours">۲۴ ساعت بعد</option>
               </select>
               <label class="flex items-center">
                 <input v-model="smsSettings.includeLink" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">شامل لینک نظرسنجی</span>
               </label>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input v-model="smsSettings.retryFailed" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">تلاش مجدد</span>
               </label>
               <label class="flex items-center">
                 <input v-model="smsSettings.unicodeSupport" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">پشتیبانی از یونیکد</span>
               </label>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Site Notifications -->
       <div class="bg-purple-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">اعلان‌های سایت</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">نمایش اعلان</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input v-model="siteNotificationSettings.enabled" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">فعال کردن اعلان‌های سایت</span>
               </label>
               <select v-model="siteNotificationSettings.position" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
                 <option value="top-right">بالا راست</option>
                 <option value="top-left">بالا چپ</option>
                 <option value="bottom-right">پایین راست</option>
                 <option value="bottom-left">پایین چپ</option>
                 <option value="center">وسط</option>
               </select>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات ظاهری</label>
             <div class="space-y-2">
               <select v-model="siteNotificationSettings.style" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
                 <option value="default">پیش‌فرض</option>
                 <option value="modern">مدرن</option>
                 <option value="minimal">مینیمال</option>
                 <option value="custom">سفارشی</option>
               </select>
               <input v-model.number="siteNotificationSettings.duration" type="number" min="1" max="30" placeholder="مدت نمایش (ثانیه)" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">محتوا</label>
             <div class="space-y-2">
               <input v-model="siteNotificationSettings.title" type="text" placeholder="عنوان اعلان" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500">
               <textarea v-model="siteNotificationSettings.message" rows="3" placeholder="پیام اعلان" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500"></textarea>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">تنظیمات پیشرفته</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input v-model="siteNotificationSettings.sound" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">صدای اعلان</span>
               </label>
               <label class="flex items-center">
                 <input v-model="siteNotificationSettings.vibration" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">لرزش (موبایل)</span>
               </label>
               <label class="flex items-center">
                 <input v-model="siteNotificationSettings.autoClose" type="checkbox" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">بستن خودکار</span>
               </label>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Notification Templates -->
       <div class="bg-gray-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">قالب‌های اطلاع‌رسانی</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
           <div v-for="template in notificationTemplates" :key="template.id" class="bg-white rounded-lg p-6 border border-gray-200">
             <div class="flex items-center justify-between mb-3">
               <h4 class="font-medium text-gray-900">{{ template.name }}</h4>
               <span :class="getTemplateStatusClass(template.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                 {{ getTemplateStatusText(template.status) }}
               </span>
             </div>
             <p class="text-sm text-gray-600 mb-3">{{ template.description }}</p>
             <div class="flex space-x-2 space-x-reverse">
               <button class="px-3 py-1 bg-blue-600 text-white text-sm rounded hover:bg-blue-700" @click="editTemplate(template)">
                 ویرایش
               </button>
               <button class="px-3 py-1 bg-gray-600 text-white text-sm rounded hover:bg-gray-700" @click="previewTemplate(template)">
                 پیش‌نمایش
               </button>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Action Buttons -->
       <div class="flex justify-between">
         <button class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="testNotifications">
           تست اطلاع‌رسانی
         </button>
         
         <div class="flex space-x-3 space-x-reverse">
           <button class="px-6 py-3 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors" @click="resetSettings">
             بازنشانی
           </button>
           <button class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" @click="saveSettings">
             ذخیره تنظیمات
           </button>
         </div>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { ref } from 'vue'
   
   interface NotificationStats {
     emailsSent: number
     emailDeliveryRate: number
     smsSent: number
     smsDeliveryRate: number
     siteNotifications: number
     siteNotificationRate: number
     responseRate: number
   }
   
   interface EmailSettings {
     autoSend: boolean
     timing: string
     template: string
     includeReward: boolean
     subject: string
     body: string
     trackOpens: boolean
     trackClicks: boolean
     retryFailed: boolean
   }
   
   interface SMSSettings {
     enabled: boolean
     provider: string
     message: string
     timing: string
     includeLink: boolean
     retryFailed: boolean
     unicodeSupport: boolean
   }
   
   interface SiteNotificationSettings {
     enabled: boolean
     position: string
     style: string
     duration: number
     title: string
     message: string
     sound: boolean
     vibration: boolean
     autoClose: boolean
   }
   
   interface NotificationTemplate {
     id: number
     name: string
     description: string
     status: 'active' | 'inactive' | 'draft'
     type: 'email' | 'sms' | 'site'
   }
   
   const notificationStats = ref<NotificationStats>({
     emailsSent: 1247,
     emailDeliveryRate: 98.5,
     smsSent: 856,
     smsDeliveryRate: 95.2,
     siteNotifications: 2341,
     siteNotificationRate: 87.3,
     responseRate: 78.9
   })
   
   const emailSettings = ref<EmailSettings>({
     autoSend: true,
     timing: '24_hours',
     template: 'default',
     includeReward: true,
     subject: 'نظرسنجی پس از خرید - {{order_number}}',
     body: 'سلام {{customer_name}}،\n\nاز خرید شما متشکریم. لطفاً نظرات خود را با ما در میان بگذارید.\n\n{{survey_link}}',
     trackOpens: true,
     trackClicks: true,
     retryFailed: true
   })
   
   const smsSettings = ref<SMSSettings>({
     enabled: true,
     provider: 'kavenegar',
     message: 'سلام {{customer_name}}، از خرید شما متشکریم. لطفاً نظرات خود را ثبت کنید: {{survey_link}}',
     timing: '24_hours',
     includeLink: true,
     retryFailed: true,
     unicodeSupport: true
   })
   
   const siteNotificationSettings = ref<SiteNotificationSettings>({
     enabled: true,
     position: 'top-right',
     style: 'modern',
     duration: 5,
     title: 'نظرسنجی جدید',
     message: 'لطفاً نظرات خود را درباره خرید اخیر ثبت کنید',
     sound: true,
     vibration: true,
     autoClose: true
   })
   
   const notificationTemplates = ref<NotificationTemplate[]>([
     {
       id: 1,
       name: 'قالب پیش‌فرض ایمیل',
       description: 'قالب استاندارد برای ارسال ایمیل نظرسنجی',
       status: 'active',
       type: 'email'
     },
     {
       id: 2,
       name: 'قالب پیامک کوتاه',
       description: 'قالب بهینه شده برای پیامک',
       status: 'active',
       type: 'sms'
     },
     {
       id: 3,
       name: 'اعلان سایت مدرن',
       description: 'قالب مدرن برای اعلان‌های سایت',
       status: 'active',
       type: 'site'
     },
     {
       id: 4,
       name: 'قالب برندینگ',
       description: 'قالب سفارشی با لوگو و رنگ‌های برند',
       status: 'draft',
       type: 'email'
     }
   ])
   
   // Methods
   const getTemplateStatusClass = (status: string) => {
     switch (status) {
       case 'active':
         return 'bg-green-100 text-green-800'
       case 'inactive':
         return 'bg-red-100 text-red-800'
       case 'draft':
         return 'bg-yellow-100 text-yellow-800'
       default:
         return 'bg-gray-100 text-gray-800'
     }
   }
   
   const getTemplateStatusText = (status: string) => {
     switch (status) {
       case 'active':
         return 'فعال'
       case 'inactive':
         return 'غیرفعال'
       case 'draft':
         return 'پیش‌نویس'
       default:
         return 'نامشخص'
     }
   }
   
   const editTemplate = (_template: NotificationTemplate) => {
     // console.log('Editing template:', template)
   }
   
   const previewTemplate = (_template: NotificationTemplate) => {
     // console.log('Previewing template:', template)
   }
   
   const testNotifications = () => {
     // console.log('Testing notifications...')
   }
   
   const saveSettings = () => {
     // console.log('Saving notification settings...')
   }
   
   const resetSettings = () => {
     if (confirm('آیا از بازنشانی تنظیمات اطلاع‌رسانی اطمینان دارید؟')) {
       // console.log('Resetting notification settings')
     }
   }
   
   // Expose methods for parent component
   defineExpose({
     emailSettings,
     smsSettings,
     siteNotificationSettings,
     notificationStats
   })
   </script>
   
   <style scoped>
   .notification-system {
     transition: all 0.3s ease;
   }
   
   /* Hover effects */
   .hover\:bg-gray-50:hover {
     background-color: #f9fafb;
   }
   
   /* Status badge animations */
   .bg-green-100,
   .bg-red-100,
   .bg-yellow-100,
   .bg-gray-100 {
     transition: all 0.2s ease;
   }
   
   /* Button hover effects */
   button:hover {
     transform: translateY(-1px);
     transition: transform 0.2s ease;
   }
   
   /* Character counter animation */
   .text-xs {
     transition: color 0.2s ease;
   }
   
   /* Responsive adjustments */
   @media (max-width: 768px) {
     .grid {
       grid-template-columns: 1fr;
     }
   }
   </style>
