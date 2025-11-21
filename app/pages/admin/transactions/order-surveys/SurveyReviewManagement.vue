<template>
     <div class="survey-review-management bg-white rounded-2xl shadow-xl border border-gray-100 p-8">
       <div class="flex items-center justify-between mb-8">
         <div class="flex items-center space-x-3 space-x-reverse">
           <div class="p-3 bg-gradient-to-r from-green-400 to-green-600 rounded-xl shadow-lg">
             <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
             </svg>
           </div>
           <div>
             <h2 class="text-2xl font-bold text-gray-900">مدیریت نظرات</h2>
             <p class="text-gray-600 mt-1">بررسی و مدیریت نظرات دریافتی از نظرسنجی‌ها</p>
           </div>
         </div>
         
         <div class="flex space-x-3 space-x-reverse">
           <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors flex items-center space-x-2 space-x-reverse" @click="exportData">
             <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
               <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
             </svg>
             <span>خروجی اکسل</span>
           </button>
           
           <button class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors" @click="bulkAction">
             عملیات گروهی
           </button>
         </div>
       </div>
   
       <!-- Filters -->
       <div class="bg-gray-50 rounded-xl p-6 mb-6">
         <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">وضعیت</label>
             <select v-model="filters.status" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
               <option value="">همه</option>
               <option value="pending">در انتظار بررسی</option>
               <option value="approved">تأیید شده</option>
               <option value="rejected">رد شده</option>
               <option value="flagged">مشکوک</option>
             </select>
           </div>
           
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">امتیاز</label>
             <select v-model="filters.rating" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
               <option value="">همه</option>
               <option value="5">۵ ستاره</option>
               <option value="4">۴ ستاره</option>
               <option value="3">۳ ستاره</option>
               <option value="2">۲ ستاره</option>
               <option value="1">۱ ستاره</option>
             </select>
           </div>
           
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">تاریخ</label>
             <select v-model="filters.dateRange" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500">
               <option value="">همه</option>
               <option value="today">امروز</option>
               <option value="week">هفته گذشته</option>
               <option value="month">ماه گذشته</option>
               <option value="quarter">سه ماه گذشته</option>
             </select>
           </div>
           
           <div>
             <label class="block text-sm font-medium text-gray-700 mb-2">جستجو</label>
             <input 
               v-model="filters.search" 
               type="text" 
               placeholder="جستجو در نظرات..."
               class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-green-500"
             >
           </div>
         </div>
         
         <div class="flex justify-between items-center mt-4">
           <div class="text-sm text-gray-600">
             {{ filteredReviews.length }} نظر از {{ totalReviews }} نظر
           </div>
           
           <div class="flex space-x-2 space-x-reverse">
             <button class="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors" @click="clearFilters">
               پاک کردن فیلترها
             </button>
             <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors" @click="applyFilters">
               اعمال فیلترها
             </button>
           </div>
         </div>
       </div>
   
       <!-- Reviews Table -->
       <div class="overflow-x-auto">
         <table class="w-full">
           <thead class="bg-gray-50">
             <tr>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                 <input v-model="selectAll" type="checkbox" class="rounded border-gray-300" @change="toggleSelectAll">
               </th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">مشتری</th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">سفارش</th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">امتیاز</th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">نظر</th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">وضعیت</th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">تاریخ</th>
               <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">عملیات</th>
             </tr>
           </thead>
           <tbody class="bg-white divide-y divide-gray-200">
             <tr v-for="review in paginatedReviews" :key="review.id" class="hover:bg-gray-50">
               <td class="px-6 py-4 whitespace-nowrap">
                 <input v-model="selectedReviews" type="checkbox" :value="review.id" class="rounded border-gray-300">
               </td>
               
               <td class="px-6 py-4 whitespace-nowrap">
                 <div class="flex items-center">
                   <div class="">
                     <img class="h-10 w-10 rounded-full" :src="review.customer.avatar" :alt="review.customer.name">
                   </div>
                   <div class="mr-4">
                     <div class="text-sm font-medium text-gray-900">{{ review.customer.name }}</div>
                     <div class="text-sm text-gray-500">{{ review.customer.email }}</div>
                   </div>
                 </div>
               </td>
               
               <td class="px-6 py-4 whitespace-nowrap">
                 <div class="text-sm text-gray-900">{{ review.orderNumber }}</div>
                 <div class="text-sm text-gray-500">{{ review.orderDate }}</div>
               </td>
               
               <td class="px-6 py-4 whitespace-nowrap">
                 <div class="flex items-center">
                   <div class="flex space-x-1 space-x-reverse">
                     <span v-for="i in 5" :key="i" class="text-yellow-400">
                       {{ i <= review.overallRating ? '★' : '☆' }}
                     </span>
                   </div>
                   <span class="text-sm text-gray-500 mr-2">({{ review.overallRating }}/5)</span>
                 </div>
               </td>
               
               <td class="px-6 py-4">
                 <div class="text-sm text-gray-900 max-w-xs truncate">
                   {{ review.comments.overallExperience || 'بدون نظر' }}
                 </div>
                 <button class="text-sm text-blue-600 hover:text-blue-800" @click="viewFullReview(review)">
                   مشاهده کامل
                 </button>
               </td>
               
               <td class="px-6 py-4 whitespace-nowrap">
                 <span :class="getStatusClass(review.status)" class="px-2 py-1 text-xs font-medium rounded-full">
                   {{ getStatusText(review.status) }}
                 </span>
               </td>
               
               <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                 {{ formatDate(review.createdAt) }}
               </td>
               
               <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                 <div class="flex space-x-2 space-x-reverse">
                   <button class="text-green-600 hover:text-green-900" @click="approveReview(review.id)">
                     تأیید
                   </button>
                   <button class="text-red-600 hover:text-red-900" @click="rejectReview(review.id)">
                     رد
                   </button>
                   <button class="text-yellow-600 hover:text-yellow-900" @click="flagReview(review.id)">
                     علامت‌گذاری
                   </button>
                   <button class="text-gray-600 hover:text-gray-900" @click="deleteReview(review.id)">
                     حذف
                   </button>
                 </div>
               </td>
             </tr>
           </tbody>
         </table>
       </div>
   
       <!-- Pagination -->
       <div class="flex items-center justify-between mt-6">
         <div class="flex items-center space-x-2 space-x-reverse">
           <span class="text-sm text-gray-700">نمایش</span>
           <select v-model="perPage" class="px-2 py-1 border border-gray-300 rounded text-sm">
             <option value="10">۱۰</option>
             <option value="25">۲۵</option>
             <option value="50">۵۰</option>
             <option value="100">۱۰۰</option>
           </select>
           <span class="text-sm text-gray-700">از {{ totalReviews }} نظر</span>
         </div>
         
         <div class="flex space-x-2 space-x-reverse">
           <button 
             :disabled="currentPage === 1" 
             class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
             @click="previousPage"
           >
             قبلی
           </button>
           
           <span class="px-3 py-1 text-sm text-gray-700">
             صفحه {{ currentPage }} از {{ totalPages }}
           </span>
           
           <button 
             :disabled="currentPage === totalPages" 
             class="px-3 py-1 border border-gray-300 rounded text-sm disabled:opacity-50"
             @click="nextPage"
           >
             بعدی
           </button>
         </div>
       </div>
   
       <!-- Review Detail Modal -->
       <div v-if="showReviewModal" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
         <div class="relative top-20 mx-auto p-5 border w-11/12 md:w-3/4 lg:w-1/2 shadow-lg rounded-md bg-white">
           <div class="mt-3">
             <div class="flex items-center justify-between mb-4">
               <h3 class="text-lg font-medium text-gray-900">جزئیات نظر</h3>
               <button class="text-gray-400 hover:text-gray-600" @click="closeReviewModal">
                 <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                 </svg>
               </button>
             </div>
             
             <div v-if="selectedReview" class="space-y-4">
               <!-- Customer Info -->
               <div class="bg-gray-50 rounded-lg p-6">
                 <h4 class="font-medium text-gray-900 mb-2">اطلاعات مشتری</h4>
                 <div class="grid grid-cols-2 gap-6 text-sm">
                   <div>
                     <span class="text-gray-600">نام:</span>
                     <span class="text-gray-900 mr-2">{{ selectedReview.customer.name }}</span>
                   </div>
                   <div>
                     <span class="text-gray-600">ایمیل:</span>
                     <span class="text-gray-900 mr-2">{{ selectedReview.customer.email }}</span>
                   </div>
                 </div>
               </div>
               
               <!-- Order Info -->
               <div class="bg-gray-50 rounded-lg p-6">
                 <h4 class="font-medium text-gray-900 mb-2">اطلاعات سفارش</h4>
                 <div class="grid grid-cols-2 gap-6 text-sm">
                   <div>
                     <span class="text-gray-600">شماره سفارش:</span>
                     <span class="text-gray-900 mr-2">{{ selectedReview.orderNumber }}</span>
                   </div>
                   <div>
                     <span class="text-gray-600">تاریخ سفارش:</span>
                     <span class="text-gray-900 mr-2">{{ selectedReview.orderDate }}</span>
                   </div>
                 </div>
               </div>
               
               <!-- Ratings -->
               <div class="bg-gray-50 rounded-lg p-6">
                 <h4 class="font-medium text-gray-900 mb-2">امتیازات</h4>
                 <div class="space-y-2">
                   <div class="flex justify-between">
                     <span class="text-sm text-gray-600">امتیاز کلی:</span>
                     <div class="flex items-center">
                       <div class="flex space-x-1 space-x-reverse">
                         <span v-for="i in 5" :key="i" class="text-yellow-400">
                           {{ i <= selectedReview.overallRating ? '★' : '☆' }}
                         </span>
                       </div>
                       <span class="text-sm text-gray-900 mr-2">({{ selectedReview.overallRating }}/5)</span>
                     </div>
                   </div>
                   
                   <div v-for="(rating, key) in selectedReview.detailedRatings" :key="key" class="flex justify-between">
                     <span class="text-sm text-gray-600">{{ getCategoryName(key) }}:</span>
                     <span class="text-sm text-gray-900">{{ rating }}/5</span>
                   </div>
                 </div>
               </div>
               
               <!-- Comments -->
               <div class="bg-gray-50 rounded-lg p-6">
                 <h4 class="font-medium text-gray-900 mb-2">نظرات</h4>
                 <div class="space-y-3">
                   <template v-for="(comment, key) in (selectedReview?.comments || {})" :key="key">
                     <div v-if="comment">
                       <div class="text-sm font-medium text-gray-700">{{ getCommentLabel(key) }}:</div>
                       <div class="text-sm text-gray-900 mt-1">{{ comment }}</div>
                     </div>
                   </template>
                 </div>
               </div>
             </div>
             
             <div class="flex justify-end space-x-3 space-x-reverse mt-6">
               <button class="px-4 py-2 bg-gray-300 text-gray-700 rounded-lg hover:bg-gray-400" @click="closeReviewModal">
                 بستن
               </button>
               <button class="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700" @click="approveReview(selectedReview?.id)">
                 تأیید
               </button>
             </div>
           </div>
         </div>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { computed, ref, watch } from 'vue'
   
   interface Customer {
     id: number
     name: string
     email: string
     avatar: string
   }
   
   interface Review {
     id: number
     customer: Customer
     orderNumber: string
     orderDate: string
     overallRating: number
     detailedRatings: Record<string, number>
     comments: Record<string, string>
     status: 'pending' | 'approved' | 'rejected' | 'flagged'
     createdAt: Date
   }
   
   interface Filters {
     status: string
     rating: string
     dateRange: string
     search: string
   }
   
   const reviews = ref<Review[]>([
     {
       id: 1,
       customer: {
         id: 1,
         name: 'علی احمدی',
         email: 'ali@example.com',
         avatar: '/images/avatars/user1.jpg'
       },
       orderNumber: 'ORD-2024-001234',
       orderDate: '۱۴۰۲/۱۰/۱۵',
       overallRating: 5,
       detailedRatings: {
         productQuality: 5,
         pricing: 4,
         deliverySpeed: 5,
         packaging: 5,
         afterSales: 4,
         userInterface: 5
       },
       comments: {
         overallExperience: 'تجربه خرید عالی بود. محصولات با کیفیت و ارسال سریع.',
         strengths: 'کیفیت محصولات و سرعت ارسال',
         weaknesses: '',
         improvements: ''
       },
       status: 'approved',
       createdAt: new Date('2024-01-15')
     },
     // Add more sample reviews...
   ])
   
   const filters = ref<Filters>({
     status: '',
     rating: '',
     dateRange: '',
     search: ''
   })
   
   const selectedReviews = ref<number[]>([])
   const selectAll = ref(false)
   const currentPage = ref(1)
   const perPage = ref(25)
   const showReviewModal = ref(false)
   const selectedReview = ref<Review | null>(null)
   
   // Computed properties
   const filteredReviews = computed(() => {
     let filtered = [...reviews.value]
     
     if (filters.value.status) {
       filtered = filtered.filter(review => review.status === filters.value.status)
     }
     
     if (filters.value.rating) {
       filtered = filtered.filter(review => review.overallRating === parseInt(filters.value.rating))
     }
     
     if (filters.value.search) {
       const searchTerm = filters.value.search.toLowerCase()
       filtered = filtered.filter(review => 
         review.customer.name.toLowerCase().includes(searchTerm) ||
         review.orderNumber.toLowerCase().includes(searchTerm) ||
         Object.values(review.comments).some(comment => 
           comment.toLowerCase().includes(searchTerm)
         )
       )
     }
     
     return filtered
   })
   
   const totalReviews = computed(() => filteredReviews.value.length)
   
   const totalPages = computed(() => Math.ceil(totalReviews.value / perPage.value))
   
   const paginatedReviews = computed(() => {
     const start = (currentPage.value - 1) * perPage.value
     const end = start + perPage.value
     return filteredReviews.value.slice(start, end)
   })
   
   // Methods
   const getStatusClass = (status: string) => {
     switch (status) {
       case 'approved':
         return 'bg-green-100 text-green-800'
       case 'rejected':
         return 'bg-red-100 text-red-800'
       case 'flagged':
         return 'bg-yellow-100 text-yellow-800'
       default:
         return 'bg-gray-100 text-gray-800'
     }
   }
   
   const getStatusText = (status: string) => {
     switch (status) {
       case 'approved':
         return 'تأیید شده'
       case 'rejected':
         return 'رد شده'
       case 'flagged':
         return 'مشکوک'
       default:
         return 'در انتظار'
     }
   }
   
   const getCategoryName = (key: string) => {
     const names = {
       productQuality: 'کیفیت محصول',
       pricing: 'قیمت‌گذاری',
       deliverySpeed: 'سرعت تحویل',
       packaging: 'بسته‌بندی',
       afterSales: 'پشتیبانی',
       userInterface: 'رابط کاربری'
     }
     return names[key] || key
   }
   
   const getCommentLabel = (key: string) => {
     const labels = {
       overallExperience: 'تجربه کلی',
       strengths: 'نقاط قوت',
       weaknesses: 'نقاط ضعف',
       improvements: 'پیشنهادات بهبود'
     }
     return labels[key] || key
   }
   
   const formatDate = (date: Date) => {
     return new Intl.DateTimeFormat('fa-IR').format(date)
   }
   
   const toggleSelectAll = () => {
     if (selectAll.value) {
       selectedReviews.value = paginatedReviews.value.map(review => review.id)
     } else {
       selectedReviews.value = []
     }
   }
   
   const viewFullReview = (review: Review) => {
     selectedReview.value = review
     showReviewModal.value = true
   }
   
   const closeReviewModal = () => {
     showReviewModal.value = false
     selectedReview.value = null
   }
   
   const approveReview = (reviewId: number) => {
     const review = reviews.value.find(r => r.id === reviewId)
     if (review) {
       review.status = 'approved'
     }
   }
   
   const rejectReview = (reviewId: number) => {
     const review = reviews.value.find(r => r.id === reviewId)
     if (review) {
       review.status = 'rejected'
     }
   }
   
   const flagReview = (reviewId: number) => {
     const review = reviews.value.find(r => r.id === reviewId)
     if (review) {
       review.status = 'flagged'
     }
   }
   
   const deleteReview = (reviewId: number) => {
     if (confirm('آیا از حذف این نظر اطمینان دارید؟')) {
       reviews.value = reviews.value.filter(r => r.id !== reviewId)
     }
   }
   
   const exportData = () => {
     // Implement export functionality
   }
   
   const bulkAction = () => {
     if (selectedReviews.value.length === 0) {
       alert('لطفاً حداقل یک نظر را انتخاب کنید')
       return
     }
     // Implement bulk actions
   }
   
   const clearFilters = () => {
     filters.value = {
       status: '',
       rating: '',
       dateRange: '',
       search: ''
     }
   }
   
   const applyFilters = () => {
     currentPage.value = 1
   }
   
   const previousPage = () => {
     if (currentPage.value > 1) {
       currentPage.value--
     }
   }
   
   const nextPage = () => {
     if (currentPage.value < totalPages.value) {
       currentPage.value++
     }
   }
   
   // Watch for changes
   watch(selectedReviews, (newValue) => {
     selectAll.value = newValue.length === paginatedReviews.value.length && newValue.length > 0
   })
   
   watch(perPage, () => {
     currentPage.value = 1
   })
   </script>
   
   <style scoped>
   .survey-review-management {
     transition: all 0.3s ease;
   }
   
   /* Table hover effects */
   tbody tr:hover {
     background-color: #f9fafb;
   }
   
   /* Modal animations */
   .modal-enter-active,
   .modal-leave-active {
     transition: opacity 0.3s ease;
   }
   
   .modal-enter-from,
   .modal-leave-to {
     opacity: 0;
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
   
   /* Responsive table */
   @media (max-width: 768px) {
     .overflow-x-auto {
       overflow-x: auto;
     }
   }
   </style>
