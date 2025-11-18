<template>
     <div class="customer-satisfaction">
       <div class="section-header">
         <h3>رضایت مشتری</h3>
         <p>تحلیل رضایت مشتریان از روش‌های ارسال</p>
       </div>
   
       <!-- Date Range Selector -->
       <div class="date-range-selector">
         <div class="form-group">
           <label>بازه زمانی:</label>
           <select v-model="selectedPeriod" @change="updateSatisfactionData">
             <option value="7">7 روز گذشته</option>
             <option value="30">30 روز گذشته</option>
             <option value="90">90 روز گذشته</option>
             <option value="365">یک سال گذشته</option>
             <option value="custom">سفارشی</option>
           </select>
         </div>
         
         <div v-if="selectedPeriod === 'custom'" class="custom-date-range">
           <div class="form-group">
             <label>از تاریخ:</label>
             <input type="date" v-model="customStartDate" @change="updateSatisfactionData">
           </div>
           <div class="form-group">
             <label>تا تاریخ:</label>
             <input type="date" v-model="customEndDate" @change="updateSatisfactionData">
           </div>
         </div>
       </div>
   
       <!-- Satisfaction Summary Cards -->
       <div class="satisfaction-summary">
         <div class="card overall-card">
           <div class="card-icon">
             <i class="fas fa-star"></i>
           </div>
           <div class="card-content">
             <h4>رضایت کلی</h4>
             <div class="rating">
               <div class="stars">
                 <i v-for="star in 5" :key="star" 
                    :class="star <= Math.round(satisfactionData.overallRating) ? 'fas fa-star filled' : 'far fa-star'"></i>
               </div>
               <div class="rating-number">{{ satisfactionData.overallRating.toFixed(1) }}/5</div>
             </div>
             <div class="trend" :class="satisfactionData.ratingChange > 0 ? 'positive' : 'negative'">
               {{ satisfactionData.ratingChange > 0 ? '+' : '' }}{{ satisfactionData.ratingChange }}%
             </div>
           </div>
         </div>
   
         <div class="card responses-card">
           <div class="card-icon">
             <i class="fas fa-comments"></i>
           </div>
           <div class="card-content">
             <h4>تعداد نظرات</h4>
             <div class="number">{{ formatNumber(satisfactionData.totalResponses) }}</div>
             <div class="trend" :class="satisfactionData.responseChange > 0 ? 'positive' : 'negative'">
               {{ satisfactionData.responseChange > 0 ? '+' : '' }}{{ satisfactionData.responseChange }}%
             </div>
           </div>
         </div>
   
         <div class="card satisfaction-card">
           <div class="card-icon">
             <i class="fas fa-smile"></i>
           </div>
           <div class="card-content">
             <h4>مشتریان راضی</h4>
             <div class="number">{{ satisfactionData.satisfiedCustomers }}%</div>
             <div class="trend" :class="satisfactionData.satisfactionChange > 0 ? 'positive' : 'negative'">
               {{ satisfactionData.satisfactionChange > 0 ? '+' : '' }}{{ satisfactionData.satisfactionChange }}%
             </div>
           </div>
         </div>
   
         <div class="card nps-card">
           <div class="card-icon">
             <i class="fas fa-thumbs-up"></i>
           </div>
           <div class="card-content">
             <h4>امتیاز NPS</h4>
             <div class="number">{{ satisfactionData.npsScore }}</div>
             <div class="trend" :class="satisfactionData.npsChange > 0 ? 'positive' : 'negative'">
               {{ satisfactionData.npsChange > 0 ? '+' : '' }}{{ satisfactionData.npsChange }}
             </div>
           </div>
         </div>
       </div>
   
       <!-- Rating Distribution -->
       <div class="rating-distribution">
         <h4>توزیع امتیازات</h4>
         <div class="distribution-chart">
           <div class="rating-bars">
             <div v-for="rating in ratingDistribution" :key="rating.stars" class="rating-bar">
               <div class="rating-label">
                 <span class="stars">
                   <i v-for="star in rating.stars" :key="star" class="fas fa-star filled"></i>
                 </span>
                 <span class="count">({{ rating.count }})</span>
               </div>
               <div class="bar-container">
                 <div class="bar" :style="{ width: rating.percentage + '%' }"></div>
                 <span class="percentage">{{ rating.percentage }}%</span>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Satisfaction by Shipping Method -->
       <div class="method-satisfaction">
         <h4>رضایت بر اساس روش ارسال</h4>
         <div class="method-grid">
           <div class="method-card" v-for="method in methodSatisfaction" :key="method.id">
             <div class="method-header">
               <h5>{{ method.name }}</h5>
               <span class="method-type">{{ method.type }}</span>
             </div>
             <div class="method-rating">
               <div class="stars">
                 <i v-for="star in 5" :key="star" 
                    :class="star <= Math.round(method.rating) ? 'fas fa-star filled' : 'far fa-star'"></i>
               </div>
               <div class="rating-number">{{ method.rating.toFixed(1) }}/5</div>
             </div>
             <div class="method-stats">
               <div class="stat">
                 <span class="label">تعداد نظرات:</span>
                 <span class="value">{{ formatNumber(method.responseCount) }}</span>
               </div>
               <div class="stat">
                 <span class="label">رضایت:</span>
                 <span class="value">{{ method.satisfactionRate }}%</span>
               </div>
               <div class="stat">
                 <span class="label">تغییر:</span>
                 <span class="value" :class="method.change >= 0 ? 'positive' : 'negative'">
                   {{ method.change >= 0 ? '+' : '' }}{{ method.change }}%
                 </span>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Feedback Analysis -->
       <div class="feedback-analysis">
         <h4>تحلیل بازخورد</h4>
         <div class="feedback-grid">
           <div class="feedback-chart">
             <h5>کلمات کلیدی مثبت</h5>
             <div class="word-cloud positive">
               <span v-for="word in positiveKeywords" :key="word.text" 
                     :style="{ fontSize: word.size + 'px' }" class="keyword">
                 {{ word.text }}
               </span>
             </div>
           </div>
   
           <div class="feedback-chart">
             <h5>کلمات کلیدی منفی</h5>
             <div class="word-cloud negative">
               <span v-for="word in negativeKeywords" :key="word.text" 
                     :style="{ fontSize: word.size + 'px' }" class="keyword">
                 {{ word.text }}
               </span>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Recent Feedback -->
       <div class="recent-feedback">
         <h4>آخرین نظرات</h4>
         <div class="feedback-list">
           <div v-for="feedback in recentFeedback" :key="feedback.id" class="feedback-item">
             <div class="feedback-header">
               <div class="customer-info">
                 <span class="customer-name">{{ feedback.customerName }}</span>
                 <span class="order-id">سفارش: {{ feedback.orderId }}</span>
               </div>
               <div class="feedback-rating">
                 <div class="stars">
                   <i v-for="star in 5" :key="star" 
                      :class="star <= feedback.rating ? 'fas fa-star filled' : 'far fa-star'"></i>
                 </div>
                 <span class="rating-text">{{ feedback.rating }}/5</span>
               </div>
             </div>
             <div class="feedback-content">
               <p>{{ feedback.comment }}</p>
             </div>
             <div class="feedback-footer">
               <span class="shipping-method">{{ feedback.shippingMethod }}</span>
               <span class="date">{{ formatDate(feedback.date) }}</span>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Export Options -->
       <div class="export-section">
         <h4>خروجی گزارش رضایت</h4>
         <div class="export-buttons">
           <button @click="exportToExcel" class="btn btn-secondary">
             <i class="fas fa-file-excel"></i>
             خروجی اکسل
           </button>
           <button @click="exportToPDF" class="btn btn-secondary">
             <i class="fas fa-file-pdf"></i>
             خروجی PDF
           </button>
           <button @click="printReport" class="btn btn-secondary">
             <i class="fas fa-print"></i>
             چاپ گزارش
           </button>
         </div>
       </div>
     </div>
   </template>
   
   <script setup lang="ts">
   import { ref, reactive, onMounted } from 'vue'
   
   // Props
   interface Props {
     shippingMethods?: any[]
   }
   
   const props = defineProps<Props>()
   
   // Emits
   const emit = defineEmits<{
     exportData: [data: any, format: string]
   }>()
   
   // Reactive data
   const selectedPeriod = ref('30')
   const customStartDate = ref('')
   const customEndDate = ref('')
   
   const satisfactionData = reactive({
     overallRating: 0,
     ratingChange: 0,
     totalResponses: 0,
     responseChange: 0,
     satisfiedCustomers: 0,
     satisfactionChange: 0,
     npsScore: 0,
     npsChange: 0
   })
   
   const ratingDistribution = ref([
     { stars: 5, count: 0, percentage: 0 },
     { stars: 4, count: 0, percentage: 0 },
     { stars: 3, count: 0, percentage: 0 },
     { stars: 2, count: 0, percentage: 0 },
     { stars: 1, count: 0, percentage: 0 }
   ])
   
   const methodSatisfaction = ref([
     {
       id: 1,
       name: 'ارسال استاندارد',
       type: 'زمینی',
       rating: 4.2,
       responseCount: 1250,
       satisfactionRate: 85,
       change: 2.5
     },
     {
       id: 2,
       name: 'ارسال سریع',
       type: 'هوایی',
       rating: 4.5,
       responseCount: 890,
       satisfactionRate: 92,
       change: 1.8
     },
     {
       id: 3,
       name: 'ارسال رایگان',
       type: 'زمینی',
       rating: 4.0,
       responseCount: 620,
       satisfactionRate: 78,
       change: -0.5
     }
   ])
   
   const positiveKeywords = ref([
     { text: 'سریع', size: 24 },
     { text: 'مطمئن', size: 20 },
     { text: 'کیفیت', size: 18 },
     { text: 'قیمت مناسب', size: 16 },
     { text: 'خدمات خوب', size: 14 },
     { text: 'به موقع', size: 12 }
   ])
   
   const negativeKeywords = ref([
     { text: 'تاخیر', size: 20 },
     { text: 'گران', size: 18 },
     { text: 'آسیب', size: 16 },
     { text: 'کند', size: 14 },
     { text: 'مشکل', size: 12 }
   ])
   
   const recentFeedback = ref([
     {
       id: 1,
       customerName: 'علی احمدی',
       orderId: 'ORD-001234',
       rating: 5,
       comment: 'ارسال بسیار سریع و با کیفیت بود. بسته‌بندی عالی و محتویات سالم رسید.',
       shippingMethod: 'ارسال سریع',
       date: new Date('2024-01-15')
     },
     {
       id: 2,
       customerName: 'فاطمه محمدی',
       orderId: 'ORD-001235',
       rating: 4,
       comment: 'کلی راضی بودم ولی کمی تاخیر داشت. قیمت مناسب بود.',
       shippingMethod: 'ارسال استاندارد',
       date: new Date('2024-01-14')
     },
     {
       id: 3,
       customerName: 'محمد رضایی',
       orderId: 'ORD-001236',
       rating: 3,
       comment: 'ارسال رایگان بود ولی خیلی کند بود. کیفیت قابل قبول.',
       shippingMethod: 'ارسال رایگان',
       date: new Date('2024-01-13')
     }
   ])
   
   // Methods
   const updateSatisfactionData = async () => {
     try {
       // Simulate API call
       await new Promise(resolve => setTimeout(resolve, 1000))
       
       // Update satisfaction data
       satisfactionData.overallRating = Math.random() * 2 + 3.5 // 3.5 to 5.5
       satisfactionData.ratingChange = Math.floor(Math.random() * 10) - 5
       satisfactionData.totalResponses = Math.floor(Math.random() * 5000) + 2000
       satisfactionData.responseChange = Math.floor(Math.random() * 20) - 10
       satisfactionData.satisfiedCustomers = Math.floor(Math.random() * 30) + 70
       satisfactionData.satisfactionChange = Math.floor(Math.random() * 15) - 7
       satisfactionData.npsScore = Math.floor(Math.random() * 60) + 20
       satisfactionData.npsChange = Math.floor(Math.random() * 20) - 10
       
       // Update rating distribution
       const total = satisfactionData.totalResponses
       ratingDistribution.value[0].count = Math.floor(total * 0.45)
       ratingDistribution.value[0].percentage = 45
       ratingDistribution.value[1].count = Math.floor(total * 0.30)
       ratingDistribution.value[1].percentage = 30
       ratingDistribution.value[2].count = Math.floor(total * 0.15)
       ratingDistribution.value[2].percentage = 15
       ratingDistribution.value[3].count = Math.floor(total * 0.07)
       ratingDistribution.value[3].percentage = 7
       ratingDistribution.value[4].count = Math.floor(total * 0.03)
       ratingDistribution.value[4].percentage = 3
     } catch (error) {
       console.error('Error updating satisfaction data:', error)
     }
   }
   
   const formatNumber = (num: number): string => {
     return new Intl.NumberFormat('fa-IR').format(num)
   }
   
   const formatDate = (date: Date): string => {
     return new Intl.DateTimeFormat('fa-IR').format(date)
   }
   
   const exportToExcel = () => {
     const data = {
       satisfactionData,
       ratingDistribution: ratingDistribution.value,
       methodSatisfaction: methodSatisfaction.value,
       recentFeedback: recentFeedback.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'excel')
   }
   
   const exportToPDF = () => {
     const data = {
       satisfactionData,
       ratingDistribution: ratingDistribution.value,
       methodSatisfaction: methodSatisfaction.value,
       recentFeedback: recentFeedback.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'pdf')
   }
   
   const printReport = () => {
     window.print()
   }
   
   // Lifecycle
   onMounted(() => {
     updateSatisfactionData()
   })
   </script>
   
   <style scoped>
   .customer-satisfaction {
     padding: 20px;
     background: #fff;
     border-radius: 8px;
     box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
   }
   
   .section-header {
     margin-bottom: 30px;
     text-align: center;
   }
   
   .section-header h3 {
     color: #2c3e50;
     margin-bottom: 10px;
     font-size: 24px;
   }
   
   .section-header p {
     color: #7f8c8d;
     font-size: 14px;
   }
   
   .date-range-selector {
     display: flex;
     gap: 20px;
     margin-bottom: 30px;
     padding: 20px;
     background: #f8f9fa;
     border-radius: 8px;
     flex-wrap: wrap;
   }
   
   .form-group {
     display: flex;
     flex-direction: column;
     gap: 5px;
   }
   
   .form-group label {
     font-weight: 600;
     color: #2c3e50;
     font-size: 14px;
   }
   
   .form-group select,
   .form-group input {
     padding: 8px 12px;
     border: 1px solid #ddd;
     border-radius: 4px;
     font-size: 14px;
   }
   
   .custom-date-range {
     display: flex;
     gap: 15px;
   }
   
   .satisfaction-summary {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
     gap: 20px;
     margin-bottom: 30px;
   }
   
   .card {
     padding: 20px;
     border-radius: 12px;
     display: flex;
     align-items: center;
     gap: 15px;
     box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
     color: white;
   }
   
   .overall-card {
     background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
   }
   
   .responses-card {
     background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
   }
   
   .satisfaction-card {
     background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
   }
   
   .nps-card {
     background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
   }
   
   .card-icon {
     font-size: 24px;
     opacity: 0.8;
   }
   
   .card-content h4 {
     margin: 0 0 10px 0;
     font-size: 14px;
     opacity: 0.9;
   }
   
   .rating {
     display: flex;
     align-items: center;
     gap: 10px;
     margin-bottom: 5px;
   }
   
   .stars {
     display: flex;
     gap: 2px;
   }
   
   .stars i {
     font-size: 16px;
   }
   
   .stars i.filled {
     color: #f39c12;
   }
   
   .rating-number {
     font-size: 18px;
     font-weight: bold;
   }
   
   .number {
     font-size: 24px;
     font-weight: bold;
     margin-bottom: 5px;
   }
   
   .trend {
     font-size: 12px;
     opacity: 0.8;
   }
   
   .trend.positive {
     color: #2ecc71;
   }
   
   .trend.negative {
     color: #e74c3c;
   }
   
   .rating-distribution {
     margin-bottom: 30px;
   }
   
   .rating-distribution h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .distribution-chart {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .rating-bars {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .rating-bar {
     display: flex;
     align-items: center;
     gap: 15px;
   }
   
   .rating-label {
     display: flex;
     align-items: center;
     gap: 8px;
     min-width: 80px;
   }
   
   .rating-label .stars {
     font-size: 12px;
   }
   
   .rating-label .count {
     font-size: 12px;
     color: #7f8c8d;
   }
   
   .bar-container {
     flex: 1;
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .bar {
     height: 20px;
     background: #3498db;
     border-radius: 10px;
     min-width: 50px;
   }
   
   .percentage {
     font-size: 12px;
     color: #2c3e50;
     min-width: 40px;
   }
   
   .method-satisfaction {
     margin-bottom: 30px;
   }
   
   .method-satisfaction h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .method-grid {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
     gap: 20px;
   }
   
   .method-card {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
     border: 1px solid #e9ecef;
   }
   
   .method-header {
     display: flex;
     justify-content: space-between;
     align-items: center;
     margin-bottom: 15px;
   }
   
   .method-header h5 {
     margin: 0;
     color: #2c3e50;
     font-size: 16px;
   }
   
   .method-type {
     background: #e9ecef;
     padding: 4px 8px;
     border-radius: 4px;
     font-size: 12px;
     color: #6c757d;
   }
   
   .method-rating {
     display: flex;
     align-items: center;
     gap: 10px;
     margin-bottom: 15px;
   }
   
   .method-stats {
     display: flex;
     flex-direction: column;
     gap: 8px;
   }
   
   .stat {
     display: flex;
     justify-content: space-between;
     font-size: 14px;
   }
   
   .stat .label {
     color: #6c757d;
   }
   
   .stat .value {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .stat .value.positive {
     color: #27ae60;
   }
   
   .stat .value.negative {
     color: #e74c3c;
   }
   
   .feedback-analysis {
     margin-bottom: 30px;
   }
   
   .feedback-analysis h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .feedback-grid {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
     gap: 30px;
   }
   
   .feedback-chart {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .feedback-chart h5 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .word-cloud {
     display: flex;
     flex-wrap: wrap;
     gap: 10px;
     justify-content: center;
     min-height: 150px;
     align-items: center;
   }
   
   .word-cloud.positive .keyword {
     color: #27ae60;
     background: rgba(39, 174, 96, 0.1);
     padding: 5px 10px;
     border-radius: 15px;
     border: 1px solid rgba(39, 174, 96, 0.3);
   }
   
   .word-cloud.negative .keyword {
     color: #e74c3c;
     background: rgba(231, 76, 60, 0.1);
     padding: 5px 10px;
     border-radius: 15px;
     border: 1px solid rgba(231, 76, 60, 0.3);
   }
   
   .recent-feedback {
     margin-bottom: 30px;
   }
   
   .recent-feedback h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .feedback-list {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .feedback-item {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
     border: 1px solid #e9ecef;
   }
   
   .feedback-header {
     display: flex;
     justify-content: space-between;
     align-items: center;
     margin-bottom: 15px;
   }
   
   .customer-info {
     display: flex;
     flex-direction: column;
     gap: 5px;
   }
   
   .customer-name {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .order-id {
     font-size: 12px;
     color: #6c757d;
   }
   
   .feedback-rating {
     display: flex;
     align-items: center;
     gap: 8px;
   }
   
   .feedback-rating .stars {
     font-size: 14px;
   }
   
   .rating-text {
     font-size: 12px;
     color: #6c757d;
   }
   
   .feedback-content {
     margin-bottom: 15px;
   }
   
   .feedback-content p {
     margin: 0;
     color: #2c3e50;
     line-height: 1.6;
   }
   
   .feedback-footer {
     display: flex;
     justify-content: space-between;
     align-items: center;
     font-size: 12px;
     color: #6c757d;
   }
   
   .shipping-method {
     background: #e9ecef;
     padding: 4px 8px;
     border-radius: 4px;
   }
   
   .export-section {
     text-align: center;
     padding: 20px;
     background: #f8f9fa;
     border-radius: 8px;
   }
   
   .export-section h4 {
     margin-bottom: 15px;
     color: #2c3e50;
   }
   
   .export-buttons {
     display: flex;
     gap: 15px;
     justify-content: center;
     flex-wrap: wrap;
   }
   
   .btn {
     padding: 10px 20px;
     border: none;
     border-radius: 6px;
     cursor: pointer;
     font-size: 14px;
     display: flex;
     align-items: center;
     gap: 8px;
     transition: all 0.3s ease;
   }
   
   .btn-secondary {
     background: #6c757d;
     color: white;
   }
   
   .btn-secondary:hover {
     background: #5a6268;
     transform: translateY(-2px);
   }
   
   @media (max-width: 768px) {
     .satisfaction-summary {
       grid-template-columns: 1fr;
     }
     
     .method-grid {
       grid-template-columns: 1fr;
     }
     
     .feedback-grid {
       grid-template-columns: 1fr;
     }
     
     .custom-date-range {
       flex-direction: column;
     }
     
     .export-buttons {
       flex-direction: column;
     }
     
     .feedback-header {
       flex-direction: column;
       align-items: flex-start;
       gap: 10px;
     }
   }
   </style>
