<template>
     <div class="security-manager bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
       <div class="flex items-center mb-6">
         <div class="p-3 bg-gradient-to-r from-red-400 to-red-600 rounded-xl shadow-lg">
           <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
             <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
           </svg>
         </div>
         <div class="mr-4">
           <h2 class="text-2xl font-bold text-gray-900">مدیریت امنیت</h2>
           <p class="text-gray-600 mt-1">تنظیمات امنیتی و محافظت از نظرسنجی‌ها</p>
         </div>
       </div>
   
       <!-- Security Status -->
       <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
         <div class="bg-gradient-to-r from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-green-600">وضعیت امنیت</p>
               <p class="text-2xl font-bold text-green-900">{{ securityStatus.level }}</p>
             </div>
             <div class="p-3 bg-green-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-green-700">
             {{ securityStatus.description }}
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-blue-600">تهدیدات مسدود شده</p>
               <p class="text-2xl font-bold text-blue-900">{{ securityStats.blockedThreats }}</p>
             </div>
             <div class="p-3 bg-blue-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728L5.636 5.636m12.728 12.728L18.364 5.636M5.636 18.364l12.728-12.728"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-blue-700">
             در ۲۴ ساعت گذشته
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-purple-600">نظرسنجی‌های تأیید شده</p>
               <p class="text-2xl font-bold text-purple-900">{{ securityStats.verifiedSurveys }}</p>
             </div>
             <div class="p-3 bg-purple-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 text-sm text-purple-700">
             {{ securityStats.verificationRate }}% نرخ تأیید
           </div>
         </div>
       </div>
   
       <!-- Anti-Spam Settings -->
       <div class="bg-gray-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">تنظیمات ضد اسپم</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">CAPTCHA</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.captcha.enabled" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">فعال کردن CAPTCHA</span>
               </label>
               <select v-model="securitySettings.captcha.type" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500">
                 <option value="recaptcha">Google reCAPTCHA</option>
                 <option value="hcaptcha">hCaptcha</option>
                 <option value="custom">CAPTCHA سفارشی</option>
               </select>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت نرخ</label>
             <div class="space-y-2">
               <div>
                 <label class="text-sm text-gray-600">حداکثر نظرسنجی در ساعت</label>
                 <input v-model.number="securitySettings.rateLimit.hourly" type="number" min="1" max="100" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500">
               </div>
               <div>
                 <label class="text-sm text-gray-600">حداکثر نظرسنجی در روز</label>
                 <input v-model.number="securitySettings.rateLimit.daily" type="number" min="1" max="1000" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500">
               </div>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">فیلتر محتوا</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.contentFilter.enabled" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">فیلتر کلمات نامناسب</span>
               </label>
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.contentFilter.blockLinks" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">مسدود کردن لینک‌ها</span>
               </label>
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.contentFilter.detectSpam" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">تشخیص خودکار اسپم</span>
               </label>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">IP Blocking</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.ipBlocking.enabled" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">مسدود کردن IP مشکوک</span>
               </label>
               <div>
                 <label class="text-sm text-gray-600">زمان مسدودیت (دقیقه)</label>
                 <input v-model.number="securitySettings.ipBlocking.duration" type="number" min="5" max="1440" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500">
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Purchase Verification -->
       <div class="bg-blue-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">تأیید خرید</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">روش تأیید</label>
             <div class="space-y-2">
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.purchaseVerification.enabled" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">تأیید خرید اجباری</span>
               </label>
               <select v-model="securitySettings.purchaseVerification.method" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
                 <option value="order_number">شماره سفارش</option>
                 <option value="email">ایمیل خریدار</option>
                 <option value="phone">شماره تلفن</option>
                 <option value="multiple">چندگانه</option>
               </select>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت زمانی</label>
             <div class="space-y-2">
               <div>
                 <label class="text-sm text-gray-600">حداکثر زمان تأیید (روز)</label>
                 <input v-model.number="securitySettings.purchaseVerification.timeLimit" type="number" min="1" max="365" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500">
               </div>
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.purchaseVerification.allowMultiple" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">اجازه نظرسنجی چندگانه</span>
               </label>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Time Limits -->
       <div class="bg-yellow-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">محدودیت‌های زمانی</h3>
         <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">زمان تکمیل نظرسنجی</label>
             <div class="space-y-2">
               <div>
                 <label class="text-sm text-gray-600">حداقل زمان (دقیقه)</label>
                 <input v-model.number="securitySettings.timeLimits.minTime" type="number" min="1" max="60" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500">
               </div>
               <div>
                 <label class="text-sm text-gray-600">حداکثر زمان (دقیقه)</label>
                 <input v-model.number="securitySettings.timeLimits.maxTime" type="number" min="5" max="120" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500">
               </div>
             </div>
           </div>
   
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">محدودیت دسترسی</label>
             <div class="space-y-2">
               <div>
                 <label class="text-sm text-gray-600">زمان انقضا لینک (روز)</label>
                 <input v-model.number="securitySettings.timeLimits.linkExpiry" type="number" min="1" max="90" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-yellow-500">
               </div>
               <label class="flex items-center">
                 <input type="checkbox" v-model="securitySettings.timeLimits.autoExpire" class="rounded border-gray-300">
                 <span class="text-sm text-gray-700 mr-2">انقضای خودکار</span>
               </label>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Security Logs -->
       <div class="bg-gray-50 rounded-xl p-6 mb-8">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">گزارش‌های امنیتی</h3>
         <div class="overflow-x-auto">
           <table class="w-full">
             <thead class="bg-gray-100">
               <tr>
                 <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">زمان</th>
                 <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">نوع</th>
                 <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">IP</th>
                 <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">عملیات</th>
                 <th class="px-4 py-2 text-right text-xs font-medium text-gray-500">وضعیت</th>
               </tr>
             </thead>
             <tbody class="bg-white divide-y divide-gray-200">
               <tr v-for="log in securityLogs" :key="log.id" class="hover:bg-gray-50">
                 <td class="px-4 py-2 text-sm text-gray-900">{{ formatTime(log.timestamp) }}</td>
                 <td class="px-4 py-2">
                   <span :class="getLogTypeClass(log.type)" class="px-2 py-1 text-xs font-medium rounded-full">
                     {{ getLogTypeText(log.type) }}
                   </span>
                 </td>
                 <td class="px-4 py-2 text-sm text-gray-600">{{ log.ip }}</td>
                 <td class="px-4 py-2 text-sm text-gray-900">{{ log.action }}</td>
                 <td class="px-4 py-2">
                   <span :class="getStatusClass(log.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                     {{ getStatusText(log.status) }}
                   </span>
                 </td>
               </tr>
             </tbody>
           </table>
         </div>
       </div>
   
       <!-- Action Buttons -->
       <div class="flex justify-between">
         <button @click="resetSettings" class="px-6 py-3 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors">
           بازنشانی تنظیمات
         </button>
         
         <div class="flex space-x-3 space-x-reverse">
           <button @click="testSecurity" class="px-6 py-3 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
             تست امنیت
           </button>
           <button @click="saveSettings" class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors">
             ذخیره تنظیمات
           </button>
         </div>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { ref } from 'vue'
   
   interface SecurityStatus {
     level: string
     description: string
   }
   
   interface SecurityStats {
     blockedThreats: number
     verifiedSurveys: number
     verificationRate: number
   }
   
   interface SecuritySettings {
     captcha: {
       enabled: boolean
       type: 'recaptcha' | 'hcaptcha' | 'custom'
     }
     rateLimit: {
       hourly: number
       daily: number
     }
     contentFilter: {
       enabled: boolean
       blockLinks: boolean
       detectSpam: boolean
     }
     ipBlocking: {
       enabled: boolean
       duration: number
     }
     purchaseVerification: {
       enabled: boolean
       method: 'order_number' | 'email' | 'phone' | 'multiple'
       timeLimit: number
       allowMultiple: boolean
     }
     timeLimits: {
       minTime: number
       maxTime: number
       linkExpiry: number
       autoExpire: boolean
     }
   }
   
   interface SecurityLog {
     id: number
     timestamp: Date
     type: 'threat' | 'verification' | 'spam' | 'block'
     ip: string
     action: string
     status: 'success' | 'blocked' | 'warning'
   }
   
   const securityStatus = ref<SecurityStatus>({
     level: 'عالی',
     description: 'تمام لایه‌های امنیتی فعال هستند'
   })
   
   const securityStats = ref<SecurityStats>({
     blockedThreats: 47,
     verifiedSurveys: 1247,
     verificationRate: 98.5
   })
   
   const securitySettings = ref<SecuritySettings>({
     captcha: {
       enabled: true,
       type: 'recaptcha'
     },
     rateLimit: {
       hourly: 5,
       daily: 50
     },
     contentFilter: {
       enabled: true,
       blockLinks: false,
       detectSpam: true
     },
     ipBlocking: {
       enabled: true,
       duration: 30
     },
     purchaseVerification: {
       enabled: true,
       method: 'order_number',
       timeLimit: 30,
       allowMultiple: false
     },
     timeLimits: {
       minTime: 2,
       maxTime: 30,
       linkExpiry: 7,
       autoExpire: true
     }
   })
   
   const securityLogs = ref<SecurityLog[]>([
     {
       id: 1,
       timestamp: new Date(Date.now() - 5 * 60 * 1000),
       type: 'threat',
       ip: '192.168.1.100',
       action: 'تلاش برای ارسال اسپم',
       status: 'blocked'
     },
     {
       id: 2,
       timestamp: new Date(Date.now() - 15 * 60 * 1000),
       type: 'verification',
       ip: '192.168.1.101',
       action: 'تأیید خرید موفق',
       status: 'success'
     },
     {
       id: 3,
       timestamp: new Date(Date.now() - 30 * 60 * 1000),
       type: 'spam',
       ip: '192.168.1.102',
       action: 'تشخیص محتوای مشکوک',
       status: 'warning'
     }
   ])
   
   // Methods
   const getLogTypeClass = (type: string) => {
     switch (type) {
       case 'threat':
         return 'bg-red-100 text-red-800'
       case 'verification':
         return 'bg-green-100 text-green-800'
       case 'spam':
         return 'bg-yellow-100 text-yellow-800'
       case 'block':
         return 'bg-gray-100 text-gray-800'
       default:
         return 'bg-gray-100 text-gray-800'
     }
   }
   
   const getLogTypeText = (type: string) => {
     switch (type) {
       case 'threat':
         return 'تهدید'
       case 'verification':
         return 'تأیید'
       case 'spam':
         return 'اسپم'
       case 'block':
         return 'مسدود'
       default:
         return 'نامشخص'
     }
   }
   
   const getStatusClass = (status: string) => {
     switch (status) {
       case 'success':
         return 'bg-green-100 text-green-800'
       case 'blocked':
         return 'bg-red-100 text-red-800'
       case 'warning':
         return 'bg-yellow-100 text-yellow-800'
       default:
         return 'bg-gray-100 text-gray-800'
     }
   }
   
   const getStatusText = (status: string) => {
     switch (status) {
       case 'success':
         return 'موفق'
       case 'blocked':
         return 'مسدود'
       case 'warning':
         return 'هشدار'
       default:
         return 'نامشخص'
     }
   }
   
   const formatTime = (timestamp: Date) => {
     return new Intl.DateTimeFormat('fa-IR', {
       hour: '2-digit',
       minute: '2-digit'
     }).format(timestamp)
   }
   
   const saveSettings = () => {
     // Save security settings to backend
     console.log('Saving security settings:', securitySettings.value)
   }
   
   const resetSettings = () => {
     if (confirm('آیا از بازنشانی تنظیمات امنیتی اطمینان دارید؟')) {
       // Reset to default settings
       console.log('Resetting security settings')
     }
   }
   
   const testSecurity = () => {
     // Test security features
     console.log('Testing security features')
   }
   
   // Expose methods for parent component
   defineExpose({
     securitySettings,
     securityStatus,
     securityStats
   })
   </script>
   
   <style scoped>
   .security-manager {
     transition: all 0.3s ease;
   }
   
   /* Hover effects */
   .hover\:bg-gray-50:hover {
     background-color: #f9fafb;
   }
   
   /* Status badge animations */
   .bg-red-100,
   .bg-green-100,
   .bg-yellow-100,
   .bg-gray-100 {
     transition: all 0.2s ease;
   }
   
   /* Button hover effects */
   button:hover {
     transform: translateY(-1px);
     transition: transform 0.2s ease;
   }
   
   /* Responsive adjustments */
   @media (max-width: 768px) {
     .grid {
       grid-template-columns: 1fr;
     }
     
     .overflow-x-auto {
       overflow-x: auto;
     }
   }
   </style>
