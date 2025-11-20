<template>
     <div class="survey-stats-dashboard bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
       <div class="flex items-center justify-between mb-8">
         <div class="flex items-center space-x-3 space-x-reverse">
           <div class="p-3 bg-gradient-to-r from-indigo-400 to-indigo-600 rounded-xl shadow-lg">
             <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
             </svg>
           </div>
           <div>
             <h2 class="text-2xl font-bold text-gray-900">آمار نظرسنجی‌ها</h2>
             <p class="text-gray-600 mt-1">نمای کلی از وضعیت نظرسنجی‌های پس از خرید</p>
           </div>
         </div>
         
         <div class="flex space-x-3 space-x-reverse">
           <select v-model="selectedPeriod" class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500">
             <option value="7">۷ روز گذشته</option>
             <option value="30">۳۰ روز گذشته</option>
             <option value="90">۹۰ روز گذشته</option>
             <option value="365">یک سال گذشته</option>
           </select>
           
           <button class="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition-colors" @click="refreshStats">
             <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"></path>
             </svg>
           </button>
         </div>
       </div>
   
       <!-- Key Metrics -->
       <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
         <div class="bg-gradient-to-r from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-blue-600">کل نظرسنجی‌ها</p>
               <p class="text-3xl font-bold text-blue-900">{{ stats.totalSurveys }}</p>
             </div>
             <div class="p-3 bg-blue-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 flex items-center">
             <span :class="stats.surveyGrowth >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
               {{ stats.surveyGrowth >= 0 ? '+' : '' }}{{ stats.surveyGrowth }}%
             </span>
             <span class="text-sm text-gray-500 mr-2">نسبت به دوره قبل</span>
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-green-50 to-green-100 rounded-xl p-6 border border-green-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-green-600">نرخ پاسخ</p>
               <p class="text-3xl font-bold text-green-900">{{ stats.responseRate }}%</p>
             </div>
             <div class="p-3 bg-green-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 flex items-center">
             <span :class="stats.responseRateGrowth >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
               {{ stats.responseRateGrowth >= 0 ? '+' : '' }}{{ stats.responseRateGrowth }}%
             </span>
             <span class="text-sm text-gray-500 mr-2">نسبت به دوره قبل</span>
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-yellow-50 to-yellow-100 rounded-xl p-6 border border-yellow-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-yellow-600">میانگین امتیاز</p>
               <p class="text-3xl font-bold text-yellow-900">{{ stats.averageRating.toFixed(1) }}</p>
             </div>
             <div class="p-3 bg-yellow-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.75l-6.172 3.247 1.179-6.873L2 9.753l6.908-1.004L12 2.25l3.092 6.499L22 9.753l-5.007 4.371 1.179 6.873z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 flex items-center">
             <div class="flex space-x-1 space-x-reverse">
               <span v-for="i in 5" :key="i" class="text-yellow-400">
                 {{ i <= Math.round(stats.averageRating) ? '★' : '☆' }}
               </span>
             </div>
           </div>
         </div>
   
         <div class="bg-gradient-to-r from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200">
           <div class="flex items-center justify-between">
             <div>
               <p class="text-sm font-medium text-purple-600">رضایت مشتری</p>
               <p class="text-3xl font-bold text-purple-900">{{ stats.customerSatisfaction }}%</p>
             </div>
             <div class="p-3 bg-purple-500 rounded-lg">
               <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.828 14.828a4 4 0 01-5.656 0M9 10h1m4 0h1m-6 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
               </svg>
             </div>
           </div>
           <div class="mt-4 flex items-center">
             <span :class="stats.satisfactionGrowth >= 0 ? 'text-green-600' : 'text-red-600'" class="text-sm font-medium">
               {{ stats.satisfactionGrowth >= 0 ? '+' : '' }}{{ stats.satisfactionGrowth }}%
             </span>
             <span class="text-sm text-gray-500 mr-2">نسبت به دوره قبل</span>
           </div>
         </div>
       </div>
   
       <!-- Charts Section -->
       <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
         <!-- Rating Distribution Chart -->
         <div class="bg-gray-50 rounded-xl p-6">
           <h3 class="text-lg font-semibold text-gray-900 mb-4">توزیع امتیازات</h3>
           <div class="space-y-3">
             <div v-for="rating in [5, 4, 3, 2, 1]" :key="rating" class="flex items-center space-x-3 space-x-reverse">
               <div class="flex items-center space-x-1 space-x-reverse w-16">
                 <span class="text-sm font-medium">{{ rating }}</span>
                 <span class="text-yellow-400">★</span>
               </div>
               <div class="flex-1 bg-gray-200 rounded-full h-3">
                 <div 
                   class="bg-yellow-400 h-3 rounded-full transition-all duration-500"
                   :style="{ width: `${getRatingPercentage(rating)}%` }"
                 ></div>
               </div>
               <span class="text-sm text-gray-600 w-12">{{ getRatingCount(rating) }}</span>
               <span class="text-sm text-gray-500 w-12">{{ getRatingPercentage(rating) }}%</span>
             </div>
           </div>
         </div>
   
         <!-- Response Trend Chart -->
         <div class="bg-gray-50 rounded-xl p-6">
           <h3 class="text-lg font-semibold text-gray-900 mb-4">روند پاسخ‌دهی</h3>
           <div class="h-48 flex items-end space-x-2 space-x-reverse">
             <div v-for="(data, index) in responseTrend" :key="index" class="flex-1 flex flex-col items-center">
               <div 
                 class="w-full bg-gradient-to-t from-blue-500 to-blue-300 rounded-t transition-all duration-300 hover:from-blue-600 hover:to-blue-400"
                 :style="{ height: `${(data.value / maxTrendValue) * 100}%` }"
               ></div>
               <span class="text-xs text-gray-600 mt-2">{{ data.label }}</span>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Recent Activity -->
       <div class="bg-gray-50 rounded-xl p-6">
         <h3 class="text-lg font-semibold text-gray-900 mb-4">فعالیت‌های اخیر</h3>
         <div class="space-y-3">
           <div v-for="activity in recentActivities" :key="activity.id" class="flex items-center space-x-3 space-x-reverse p-3 bg-white rounded-lg">
             <div class="p-2 rounded-lg" :class="getActivityIconClass(activity.type)">
               <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                 <path v-if="activity.type === 'survey'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                 <path v-else-if="activity.type === 'rating'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.75l-6.172 3.247 1.179-6.873L2 9.753l6.908-1.004L12 2.25l3.092 6.499L22 9.753l-5.007 4.371 1.179 6.873z"></path>
                 <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
               </svg>
             </div>
             <div class="flex-1">
               <p class="text-sm font-medium text-gray-900">{{ activity.title }}</p>
               <p class="text-xs text-gray-500">{{ activity.description }}</p>
             </div>
             <span class="text-xs text-gray-400">{{ formatTime(activity.timestamp) }}</span>
           </div>
         </div>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { computed, onMounted, ref, watch } from 'vue'
   
   interface SurveyStats {
     totalSurveys: number
     surveyGrowth: number
     responseRate: number
     responseRateGrowth: number
     averageRating: number
     customerSatisfaction: number
     satisfactionGrowth: number
     ratingDistribution: Record<number, number>
   }
   
   interface ResponseTrend {
     label: string
     value: number
   }
   
   interface RecentActivity {
     id: number
     type: 'survey' | 'rating' | 'comment'
     title: string
     description: string
     timestamp: Date
   }
   
   const selectedPeriod = ref('30')
   const loading = ref(false)
   
   // Sample data - in real app this would come from API
   const stats = ref<SurveyStats>({
     totalSurveys: 1247,
     surveyGrowth: 12.5,
     responseRate: 78.3,
     responseRateGrowth: 5.2,
     averageRating: 4.2,
     customerSatisfaction: 85.7,
     satisfactionGrowth: 8.1,
     ratingDistribution: {
       5: 456,
       4: 389,
       3: 234,
       2: 98,
       1: 70
     }
   })
   
   const responseTrend = ref<ResponseTrend[]>([
     { label: 'شنبه', value: 45 },
     { label: 'یکشنبه', value: 52 },
     { label: 'دوشنبه', value: 38 },
     { label: 'سه‌شنبه', value: 61 },
     { label: 'چهارشنبه', value: 48 },
     { label: 'پنج‌شنبه', value: 55 },
     { label: 'جمعه', value: 42 }
   ])
   
   const recentActivities = ref<RecentActivity[]>([
     {
       id: 1,
       type: 'survey',
       title: 'نظرسنجی جدید ثبت شد',
       description: 'سفارش #ORD-2024-001234',
       timestamp: new Date(Date.now() - 5 * 60 * 1000)
     },
     {
       id: 2,
       type: 'rating',
       title: 'امتیاز ۵ ستاره دریافت شد',
       description: 'محصول: لپ تاپ اپل مک‌بوک پرو',
       timestamp: new Date(Date.now() - 15 * 60 * 1000)
     },
     {
       id: 3,
       type: 'comment',
       title: 'نظر جدید اضافه شد',
       description: 'تجربه خرید عالی بود',
       timestamp: new Date(Date.now() - 30 * 60 * 1000)
     },
     {
       id: 4,
       type: 'survey',
       title: 'نظرسنجی جدید ثبت شد',
       description: 'سفارش #ORD-2024-001235',
       timestamp: new Date(Date.now() - 45 * 60 * 1000)
     }
   ])
   
   // Computed properties
   const maxTrendValue = computed(() => {
     return Math.max(...responseTrend.value.map(item => item.value))
   })
   
   // Methods
   const getRatingPercentage = (rating: number) => {
     const total = Object.values(stats.value.ratingDistribution).reduce((sum, count) => sum + count, 0)
     return Math.round((stats.value.ratingDistribution[rating] / total) * 100)
   }
   
   const getRatingCount = (rating: number) => {
     return stats.value.ratingDistribution[rating] || 0
   }
   
   const getActivityIconClass = (type: string) => {
     switch (type) {
       case 'survey':
         return 'bg-blue-500'
       case 'rating':
         return 'bg-yellow-500'
       case 'comment':
         return 'bg-green-500'
       default:
         return 'bg-gray-500'
     }
   }
   
   const formatTime = (timestamp: Date) => {
     const now = new Date()
     const diff = now.getTime() - timestamp.getTime()
     const minutes = Math.floor(diff / (1000 * 60))
     const hours = Math.floor(diff / (1000 * 60 * 60))
     const days = Math.floor(diff / (1000 * 60 * 60 * 24))
   
     if (minutes < 60) {
       return `${minutes} دقیقه پیش`
     } else if (hours < 24) {
       return `${hours} ساعت پیش`
     } else {
       return `${days} روز پیش`
     }
   }
   
   const refreshStats = async () => {
     loading.value = true
     try {
       // Simulate API call
       await new Promise(resolve => setTimeout(resolve, 1000))
       // Update stats here
     } catch (error) {
       console.error('Error refreshing stats:', error)
     } finally {
       loading.value = false
     }
   }
   
   // Watch for period changes
   watch(selectedPeriod, () => {
     refreshStats()
   })
   
   onMounted(() => {
     refreshStats()
   })
   
   // Expose methods for parent component
   defineExpose({
     refreshStats,
     stats
   })
   </script>
   
   <style scoped>
   .survey-stats-dashboard {
     transition: all 0.3s ease;
   }
   
   /* Smooth animations for charts */
   .bg-yellow-400 {
     transition: width 0.5s ease-out;
   }
   
   .bg-gradient-to-t {
     transition: all 0.3s ease;
   }
   
   /* Hover effects */
   .bg-gradient-to-t:hover {
     transform: scaleY(1.05);
   }
   
   /* Loading animation */
   @keyframes pulse {
     0%, 100% {
       opacity: 1;
     }
     50% {
       opacity: 0.7;
     }
   }
   
   .loading {
     animation: pulse 2s infinite;
   }
   
   /* Responsive adjustments */
   @media (max-width: 768px) {
     .grid {
       grid-template-columns: 1fr;
     }
   }
   </style>
