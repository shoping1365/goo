<template>
  <div class="issues-complaints">
       <div class="section-header">
         <h3>مشکلات و شکایات</h3>
         <p>مدیریت و تحلیل مشکلات و شکایات مشتریان</p>
       </div>
   
       <!-- Date Range Selector -->
       <div class="date-range-selector">
         <div class="form-group">
           <label>بازه زمانی:</label>
           <select v-model="selectedPeriod" @change="updateIssuesData">
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
             <input v-model="customStartDate" type="date" @change="updateIssuesData">
           </div>
           <div class="form-group">
             <label>تا تاریخ:</label>
             <input v-model="customEndDate" type="date" @change="updateIssuesData">
           </div>
         </div>
       </div>
   
       <!-- Issues Summary Cards -->
       <div class="issues-summary">
         <div class="card total-card">
           <div class="card-icon">
             <i class="fas fa-exclamation-triangle"></i>
           </div>
           <div class="card-content">
             <h4>کل مشکلات</h4>
             <div class="number">{{ formatNumber(issuesData.totalIssues) }}</div>
             <div class="trend" :class="issuesData.issuesChange > 0 ? 'negative' : 'positive'">
               {{ issuesData.issuesChange > 0 ? '+' : '' }}{{ issuesData.issuesChange }}%
             </div>
           </div>
         </div>
   
         <div class="card resolved-card">
           <div class="card-icon">
             <i class="fas fa-check-circle"></i>
           </div>
           <div class="card-content">
             <h4>مشکلات حل شده</h4>
             <div class="number">{{ issuesData.resolvedIssues }}%</div>
             <div class="trend" :class="issuesData.resolutionChange > 0 ? 'positive' : 'negative'">
               {{ issuesData.resolutionChange > 0 ? '+' : '' }}{{ issuesData.resolutionChange }}%
             </div>
           </div>
         </div>
   
         <div class="card pending-card">
           <div class="card-icon">
             <i class="fas fa-clock"></i>
           </div>
           <div class="card-content">
             <h4>در انتظار حل</h4>
             <div class="number">{{ formatNumber(issuesData.pendingIssues) }}</div>
             <div class="trend" :class="issuesData.pendingChange > 0 ? 'negative' : 'positive'">
               {{ issuesData.pendingChange > 0 ? '+' : '' }}{{ issuesData.pendingChange }}%
             </div>
           </div>
         </div>
   
         <div class="card satisfaction-card">
           <div class="card-icon">
             <i class="fas fa-smile"></i>
           </div>
           <div class="card-content">
             <h4>رضایت از حل مشکل</h4>
             <div class="number">{{ issuesData.resolutionSatisfaction }}%</div>
             <div class="trend" :class="issuesData.satisfactionChange > 0 ? 'positive' : 'negative'">
               {{ issuesData.satisfactionChange > 0 ? '+' : '' }}{{ issuesData.satisfactionChange }}%
             </div>
           </div>
         </div>
       </div>
   
       <!-- Issue Categories -->
       <div class="issue-categories">
         <h4>دسته‌بندی مشکلات</h4>
         <div class="categories-grid">
           <div v-for="category in issueCategories" :key="category.id" class="category-card">
             <div class="category-header">
               <div class="category-icon" :style="{ backgroundColor: category.color }">
                 <i :class="category.icon"></i>
               </div>
               <div class="category-info">
                 <h5>{{ category.name }}</h5>
                 <span class="category-count">{{ formatNumber(category.count) }} مورد</span>
               </div>
             </div>
             <div class="category-stats">
               <div class="stat">
                 <span class="label">درصد کل:</span>
                 <span class="value">{{ category.percentage }}%</span>
               </div>
               <div class="stat">
                 <span class="label">حل شده:</span>
                 <span class="value">{{ category.resolved }}%</span>
               </div>
               <div class="stat">
                 <span class="label">زمان حل:</span>
                 <span class="value">{{ category.avgResolutionTime }} ساعت</span>
               </div>
             </div>
             <div class="category-trend">
               <div class="trend-bar">
                 <div class="bar" :style="{ width: category.trend + '%', backgroundColor: category.color }"></div>
               </div>
               <span class="trend-text" :class="category.trend > 0 ? 'negative' : 'positive'">
                 {{ category.trend > 0 ? '+' : '' }}{{ category.trend }}%
               </span>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Resolution Time Analysis -->
       <div class="resolution-analysis">
         <h4>تحلیل زمان حل مشکلات</h4>
         <div class="resolution-grid">
           <div class="resolution-chart">
             <h5>زمان حل بر اساس دسته‌بندی</h5>
             <div class="chart-wrapper">
               <canvas ref="resolutionChart" width="400" height="200"></canvas>
             </div>
           </div>
   
           <div class="resolution-stats">
             <h5>آمار زمان حل</h5>
             <div class="stats-list">
               <div class="stat-item">
                 <span class="label">میانگین زمان حل:</span>
                 <span class="value">{{ resolutionStats.averageTime }} ساعت</span>
               </div>
               <div class="stat-item">
                 <span class="label">سریع‌ترین حل:</span>
                 <span class="value">{{ resolutionStats.fastestTime }} ساعت</span>
               </div>
               <div class="stat-item">
                 <span class="label">کندترین حل:</span>
                 <span class="value">{{ resolutionStats.slowestTime }} ساعت</span>
               </div>
               <div class="stat-item">
                 <span class="label">مشکلات حل شده در 24 ساعت:</span>
                 <span class="value">{{ resolutionStats.resolvedIn24h }}%</span>
               </div>
               <div class="stat-item">
                 <span class="label">مشکلات حل شده در 48 ساعت:</span>
                 <span class="value">{{ resolutionStats.resolvedIn48h }}%</span>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Customer Support Performance -->
       <div class="support-performance">
         <h4>عملکرد پشتیبانی مشتری</h4>
         <div class="support-grid">
           <div class="support-card">
             <h5>زمان پاسخگویی</h5>
             <div class="response-times">
               <div v-for="time in responseTimes" :key="time.range" class="response-item">
                 <div class="response-info">
                   <span class="range">{{ time.range }}</span>
                   <span class="count">{{ time.count }} مورد</span>
                 </div>
                 <div class="response-bar">
                   <div class="bar" :style="{ width: time.percentage + '%' }"></div>
                   <span class="percentage">{{ time.percentage }}%</span>
                 </div>
               </div>
             </div>
           </div>
   
           <div class="support-card">
             <h5>رضایت از پشتیبانی</h5>
             <div class="satisfaction-rating">
               <div class="rating-display">
                 <div class="stars">
                   <i
v-for="star in 5" :key="star" 
                      :class="star <= Math.round(supportSatisfaction.rating) ? 'fas fa-star filled' : 'far fa-star'"></i>
                 </div>
                 <div class="rating-number">{{ supportSatisfaction.rating.toFixed(1) }}/5</div>
               </div>
               <div class="rating-breakdown">
                 <div v-for="rating in supportSatisfaction.breakdown" :key="rating.stars" class="rating-item">
                   <span class="stars">
                     <i v-for="star in rating.stars" :key="star" class="fas fa-star filled"></i>
                   </span>
                   <div class="rating-bar">
                     <div class="bar" :style="{ width: rating.percentage + '%' }"></div>
                   </div>
                   <span class="percentage">{{ rating.percentage }}%</span>
                 </div>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Recent Issues -->
       <div class="recent-issues">
         <h4>آخرین مشکلات</h4>
         <div class="issues-list">
           <div v-for="issue in recentIssues" :key="issue.id" class="issue-item">
             <div class="issue-header">
               <div class="issue-info">
                 <span class="issue-id">#{{ issue.id }}</span>
                 <span class="customer-name">{{ issue.customerName }}</span>
                 <span class="order-id">سفارش: {{ issue.orderId }}</span>
               </div>
               <div class="issue-status" :class="issue.status">
                 <i :class="getStatusIcon(issue.status)"></i>
                 <span>{{ getStatusText(issue.status) }}</span>
               </div>
             </div>
             <div class="issue-content">
               <div class="issue-category">
                 <span class="label">دسته‌بندی:</span>
                 <span class="value">{{ issue.category }}</span>
               </div>
               <div class="issue-description">
                 <p>{{ issue.description }}</p>
               </div>
               <div class="issue-meta">
                 <div class="meta-item">
                   <span class="label">تاریخ ثبت:</span>
                   <span class="value">{{ formatDate(issue.createdAt) }}</span>
                 </div>
                 <div class="meta-item">
                   <span class="label">اولویت:</span>
                   <span class="value" :class="'priority-' + issue.priority">{{ getPriorityText(issue.priority) }}</span>
                 </div>
                 <div class="meta-item">
                   <span class="label">زمان حل:</span>
                   <span class="value">{{ issue.resolutionTime || '-' }}</span>
                 </div>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Export Options -->
       <div class="export-section">
         <h4>خروجی گزارش مشکلات</h4>
         <div class="export-buttons">
           <button class="btn btn-secondary" @click="exportToExcel">
             <i class="fas fa-file-excel"></i>
             خروجی اکسل
           </button>
           <button class="btn btn-secondary" @click="exportToPDF">
             <i class="fas fa-file-pdf"></i>
             خروجی PDF
           </button>
           <button class="btn btn-secondary" @click="printReport">
             <i class="fas fa-print"></i>
             چاپ گزارش
           </button>
         </div>
       </div>
     </div>
   </template>
   
   <script lang="ts">
   declare const definePageMeta: (meta: { layout?: string; middleware?: string }) => void
   // declare const useAuth: () => { user: unknown; hasPermission: (perm: string) => boolean }
   </script>
   
   <script setup lang="ts">
   import { nextTick, onMounted, reactive, ref } from 'vue';
   
definePageMeta({
  layout: 'admin-main',
  middleware: 'admin'
})

// استفاده از useAuth برای چک کردن پرمیژن‌ها
// const { user, hasPermission } = useAuth()
   
   // Props
   interface Props {
     shippingMethods?: unknown[]
   }
   
   defineProps<Props>()
   
   // Emits
   const emit = defineEmits<{
     exportData: [data: unknown, format: string]
   }>()
   
   // Reactive data
   const selectedPeriod = ref('30')
   const customStartDate = ref('')
   const customEndDate = ref('')
   
   const issuesData = reactive({
     totalIssues: 0,
     issuesChange: 0,
     resolvedIssues: 0,
     resolutionChange: 0,
     pendingIssues: 0,
     pendingChange: 0,
     resolutionSatisfaction: 0,
     satisfactionChange: 0
   })
   
   const issueCategories = ref([
     {
       id: 1,
       name: 'تاخیر در تحویل',
       icon: 'fas fa-clock',
       color: '#e74c3c',
       count: 0,
       percentage: 0,
       resolved: 0,
       avgResolutionTime: 0,
       trend: 0
     },
     {
       id: 2,
       name: 'آسیب به کالا',
       icon: 'fas fa-box-open',
       color: '#f39c12',
       count: 0,
       percentage: 0,
       resolved: 0,
       avgResolutionTime: 0,
       trend: 0
     },
     {
       id: 3,
       name: 'مشکلات بسته‌بندی',
       icon: 'fas fa-box',
       color: '#3498db',
       count: 0,
       percentage: 0,
       resolved: 0,
       avgResolutionTime: 0,
       trend: 0
     },
     {
       id: 4,
       name: 'مشکلات پرداخت',
       icon: 'fas fa-credit-card',
       color: '#9b59b6',
       count: 0,
       percentage: 0,
       resolved: 0,
       avgResolutionTime: 0,
       trend: 0
     },
     {
       id: 5,
       name: 'مشکلات ردیابی',
       icon: 'fas fa-search',
       color: '#1abc9c',
       count: 0,
       percentage: 0,
       resolved: 0,
       avgResolutionTime: 0,
       trend: 0
     },
     {
       id: 6,
       name: 'سایر مشکلات',
       icon: 'fas fa-exclamation',
       color: '#95a5a6',
       count: 0,
       percentage: 0,
       resolved: 0,
       avgResolutionTime: 0,
       trend: 0
     }
   ])
   
   const resolutionStats = reactive({
     averageTime: 0,
     fastestTime: 0,
     slowestTime: 0,
     resolvedIn24h: 0,
     resolvedIn48h: 0
   })
   
   const responseTimes = ref([
     { range: 'کمتر از 1 ساعت', count: 0, percentage: 0 },
     { range: '1-4 ساعت', count: 0, percentage: 0 },
     { range: '4-8 ساعت', count: 0, percentage: 0 },
     { range: '8-24 ساعت', count: 0, percentage: 0 },
     { range: 'بیش از 24 ساعت', count: 0, percentage: 0 }
   ])
   
   const supportSatisfaction = reactive({
     rating: 0,
     breakdown: [
       { stars: 5, percentage: 0 },
       { stars: 4, percentage: 0 },
       { stars: 3, percentage: 0 },
       { stars: 2, percentage: 0 },
       { stars: 1, percentage: 0 }
     ]
   })
   
   const recentIssues = ref([
     {
       id: 'ISS-001',
       customerName: 'علی احمدی',
       orderId: 'ORD-001234',
       status: 'resolved',
       category: 'تاخیر در تحویل',
       description: 'سفارش من با تاخیر 2 روزه تحویل داده شد. انتظار داشتم در زمان مشخص شده تحویل شود.',
       createdAt: new Date('2024-01-15'),
       priority: 'medium',
       resolutionTime: '24 ساعت'
     },
     {
       id: 'ISS-002',
       customerName: 'فاطمه محمدی',
       orderId: 'ORD-001235',
       status: 'pending',
       category: 'آسیب به کالا',
       description: 'کالای دریافتی آسیب دیده بود. بسته‌بندی پاره شده و محتویات شکسته بود.',
       createdAt: new Date('2024-01-14'),
       priority: 'high',
       resolutionTime: null
     },
     {
       id: 'ISS-003',
       customerName: 'محمد رضایی',
       orderId: 'ORD-001236',
       status: 'in_progress',
       category: 'مشکلات بسته‌بندی',
       description: 'بسته‌بندی نامناسب بود و کالا در معرض آسیب قرار داشت.',
       createdAt: new Date('2024-01-13'),
       priority: 'low',
       resolutionTime: null
     }
   ])
   
   // Chart references
   const resolutionChart = ref<HTMLCanvasElement>()
   
   // Methods
   const updateIssuesData = async () => {
     try {
       // Simulate API call
       await new Promise(resolve => setTimeout(resolve, 1000))
       
       // Update issues data
       issuesData.totalIssues = Math.floor(Math.random() * 500) + 200
       issuesData.issuesChange = Math.floor(Math.random() * 20) - 10
       issuesData.resolvedIssues = Math.floor(Math.random() * 30) + 70
       issuesData.resolutionChange = Math.floor(Math.random() * 15) - 7
       issuesData.pendingIssues = Math.floor(Math.random() * 100) + 50
       issuesData.pendingChange = Math.floor(Math.random() * 20) - 10
       issuesData.resolutionSatisfaction = Math.floor(Math.random() * 30) + 70
       issuesData.satisfactionChange = Math.floor(Math.random() * 15) - 7
       
       // Update issue categories
       const total = issuesData.totalIssues
       const percentages = [35, 25, 15, 12, 8, 5]
       const resolvedRates = [85, 90, 88, 95, 92, 80]
       const avgTimes = [12, 8, 6, 4, 10, 15]
       
       issueCategories.value.forEach((category, index) => {
         category.count = Math.floor(total * (percentages[index] / 100))
         category.percentage = percentages[index]
         category.resolved = resolvedRates[index]
         category.avgResolutionTime = avgTimes[index]
         category.trend = Math.floor(Math.random() * 20) - 10
       })
       
       // Update resolution stats
       resolutionStats.averageTime = Math.floor(Math.random() * 10) + 8
       resolutionStats.fastestTime = Math.floor(Math.random() * 4) + 2
       resolutionStats.slowestTime = Math.floor(Math.random() * 20) + 30
       resolutionStats.resolvedIn24h = Math.floor(Math.random() * 30) + 60
       resolutionStats.resolvedIn48h = Math.floor(Math.random() * 20) + 80
       
       // Update response times
       const responsePercentages = [40, 30, 20, 8, 2]
       responseTimes.value.forEach((time, index) => {
         time.count = Math.floor(total * (responsePercentages[index] / 100))
         time.percentage = responsePercentages[index]
       })
       
       // Update support satisfaction
       supportSatisfaction.rating = Math.random() * 2 + 3.5
       const breakdownPercentages = [45, 30, 15, 7, 3]
       supportSatisfaction.breakdown.forEach((rating, index) => {
         rating.percentage = breakdownPercentages[index]
       })
       
       updateCharts()
     } catch (error) {
       console.error('Error updating issues data:', error)
     }
   }
   
   const updateCharts = () => {
     nextTick(() => {
       createResolutionChart()
     })
   }
   
   const createResolutionChart = () => {
     if (!resolutionChart.value) return
     
     const ctx = resolutionChart.value.getContext('2d')
     if (!ctx) return
     
     // Simple bar chart implementation
     const canvas = resolutionChart.value
     const width = canvas.width
     const height = canvas.height
     
     ctx.clearRect(0, 0, width, height)
     
     const data = issueCategories.value.map(item => item.avgResolutionTime)
     const maxValue = Math.max(...data)
     const barWidth = width / data.length - 10
     
     data.forEach((value, index) => {
       const barHeight = (value / maxValue) * (height - 40)
       const x = index * (barWidth + 10) + 5
       const y = height - barHeight - 20
       
       ctx.fillStyle = issueCategories.value[index].color
       ctx.fillRect(x, y, barWidth, barHeight)
       
       // Draw value label
       ctx.fillStyle = '#2c3e50'
       ctx.font = '12px Arial'
       ctx.textAlign = 'center'
       ctx.fillText(value + ' ساعت', x + barWidth / 2, y - 5)
     })
   }
   
   const getStatusIcon = (status: string): string => {
     switch (status) {
       case 'resolved': return 'fas fa-check-circle'
       case 'pending': return 'fas fa-clock'
       case 'in_progress': return 'fas fa-tools'
       default: return 'fas fa-exclamation'
     }
   }
   
   const getStatusText = (status: string): string => {
     switch (status) {
       case 'resolved': return 'حل شده'
       case 'pending': return 'در انتظار'
       case 'in_progress': return 'در حال بررسی'
       default: return 'نامشخص'
     }
   }
   
   const getPriorityText = (priority: string): string => {
     switch (priority) {
       case 'high': return 'بالا'
       case 'medium': return 'متوسط'
       case 'low': return 'پایین'
       default: return 'نامشخص'
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
       issuesData,
       issueCategories: issueCategories.value,
       resolutionStats,
       responseTimes: responseTimes.value,
       supportSatisfaction,
       recentIssues: recentIssues.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'excel')
   }
   
   const exportToPDF = () => {
     const data = {
       issuesData,
       issueCategories: issueCategories.value,
       resolutionStats,
       responseTimes: responseTimes.value,
       supportSatisfaction,
       recentIssues: recentIssues.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'pdf')
   }
   
   const printReport = () => {
     window.print()
   }
   
   // Lifecycle
   onMounted(() => {
     updateIssuesData()
   })
   </script>
   
   <style scoped>
   .issues-complaints {
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
   
   .issues-summary {
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
   
   .total-card {
     background: linear-gradient(135deg, #e74c3c 0%, #c0392b 100%);
   }
   
   .resolved-card {
     background: linear-gradient(135deg, #27ae60 0%, #2ecc71 100%);
   }
   
   .pending-card {
     background: linear-gradient(135deg, #f39c12 0%, #e67e22 100%);
   }
   
   .satisfaction-card {
     background: linear-gradient(135deg, #3498db 0%, #2980b9 100%);
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
   
   .issue-categories {
     margin-bottom: 30px;
   }
   
   .issue-categories h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .categories-grid {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
     gap: 20px;
   }
   
   .category-card {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
     border: 1px solid #e9ecef;
   }
   
   .category-header {
     display: flex;
     align-items: center;
     gap: 15px;
     margin-bottom: 15px;
   }
   
   .category-icon {
     width: 50px;
     height: 50px;
     border-radius: 50%;
     display: flex;
     align-items: center;
     justify-content: center;
     color: white;
     font-size: 20px;
   }
   
   .category-info h5 {
     margin: 0 0 5px 0;
     color: #2c3e50;
     font-size: 16px;
   }
   
   .category-count {
     font-size: 12px;
     color: #6c757d;
   }
   
   .category-stats {
     display: flex;
     flex-direction: column;
     gap: 8px;
     margin-bottom: 15px;
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
   
   .category-trend {
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .trend-bar {
     flex: 1;
     height: 8px;
     background: #e9ecef;
     border-radius: 4px;
     overflow: hidden;
   }
   
   .trend-bar .bar {
     height: 100%;
     border-radius: 4px;
     transition: width 0.3s ease;
   }
   
   .trend-text {
     font-size: 12px;
     min-width: 50px;
   }
   
   .trend-text.positive {
     color: #27ae60;
   }
   
   .trend-text.negative {
     color: #e74c3c;
   }
   
   .resolution-analysis {
     margin-bottom: 30px;
   }
   
   .resolution-analysis h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .resolution-grid {
     display: grid;
     grid-template-columns: 2fr 1fr;
     gap: 30px;
   }
   
   .resolution-chart {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .resolution-chart h5 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .chart-wrapper {
     position: relative;
     height: 300px;
   }
   
   .resolution-stats {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .resolution-stats h5 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .stats-list {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .stat-item {
     display: flex;
     justify-content: space-between;
     align-items: center;
     padding: 10px;
     background: white;
     border-radius: 6px;
     border: 1px solid #e9ecef;
   }
   
   .stat-item .label {
     color: #6c757d;
     font-size: 14px;
   }
   
   .stat-item .value {
     font-weight: 600;
     color: #2c3e50;
     font-size: 14px;
   }
   
   .support-performance {
     margin-bottom: 30px;
   }
   
   .support-performance h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .support-grid {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
     gap: 30px;
   }
   
   .support-card {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .support-card h5 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .response-times {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .response-item {
     display: flex;
     flex-direction: column;
     gap: 8px;
   }
   
   .response-info {
     display: flex;
     justify-content: space-between;
     align-items: center;
     font-size: 14px;
   }
   
   .range {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .count {
     color: #6c757d;
     font-size: 12px;
   }
   
   .response-bar {
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .response-bar .bar {
     height: 12px;
     background: #3498db;
     border-radius: 6px;
     min-width: 100px;
   }
   
   .percentage {
     font-size: 12px;
     color: #2c3e50;
     min-width: 40px;
   }
   
   .satisfaction-rating {
     display: flex;
     flex-direction: column;
     gap: 20px;
   }
   
   .rating-display {
     display: flex;
     align-items: center;
     justify-content: center;
     gap: 15px;
   }
   
   .stars {
     display: flex;
     gap: 2px;
   }
   
   .stars i {
     font-size: 20px;
   }
   
   .stars i.filled {
     color: #f39c12;
   }
   
   .rating-number {
     font-size: 24px;
     font-weight: bold;
     color: #2c3e50;
   }
   
   .rating-breakdown {
     display: flex;
     flex-direction: column;
     gap: 10px;
   }
   
   .rating-item {
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .rating-item .stars {
     font-size: 12px;
     min-width: 60px;
   }
   
   .rating-bar {
     flex: 1;
     height: 8px;
     background: #e9ecef;
     border-radius: 4px;
     overflow: hidden;
   }
   
   .rating-bar .bar {
     height: 100%;
     background: #f39c12;
     border-radius: 4px;
   }
   
   .recent-issues {
     margin-bottom: 30px;
   }
   
   .recent-issues h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .issues-list {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .issue-item {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
     border: 1px solid #e9ecef;
   }
   
   .issue-header {
     display: flex;
     justify-content: space-between;
     align-items: center;
     margin-bottom: 15px;
   }
   
   .issue-info {
     display: flex;
     flex-direction: column;
     gap: 5px;
   }
   
   .issue-id {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .customer-name {
     font-size: 14px;
     color: #6c757d;
   }
   
   .order-id {
     font-size: 12px;
     color: #6c757d;
   }
   
   .issue-status {
     display: flex;
     align-items: center;
     gap: 8px;
     padding: 6px 12px;
     border-radius: 20px;
     font-size: 12px;
     font-weight: 600;
   }
   
   .issue-status.resolved {
     background: rgba(39, 174, 96, 0.1);
     color: #27ae60;
   }
   
   .issue-status.pending {
     background: rgba(243, 156, 18, 0.1);
     color: #f39c12;
   }
   
   .issue-status.in_progress {
     background: rgba(52, 152, 219, 0.1);
     color: #3498db;
   }
   
   .issue-content {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .issue-category {
     display: flex;
     gap: 10px;
     font-size: 14px;
   }
   
   .issue-category .label {
     color: #6c757d;
   }
   
   .issue-category .value {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .issue-description p {
     margin: 0;
     color: #2c3e50;
     line-height: 1.6;
     font-size: 14px;
   }
   
   .issue-meta {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
     gap: 10px;
   }
   
   .meta-item {
     display: flex;
     justify-content: space-between;
     font-size: 14px;
   }
   
   .meta-item .label {
     color: #6c757d;
   }
   
   .meta-item .value {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .priority-high {
     color: #e74c3c;
   }
   
   .priority-medium {
     color: #f39c12;
   }
   
   .priority-low {
     color: #27ae60;
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
     .issues-summary {
       grid-template-columns: 1fr;
     }
     
     .categories-grid {
       grid-template-columns: 1fr;
     }
     
     .resolution-grid {
       grid-template-columns: 1fr;
     }
     
     .support-grid {
       grid-template-columns: 1fr;
     }
     
     .custom-date-range {
       flex-direction: column;
     }
     
     .export-buttons {
       flex-direction: column;
     }
     
     .issue-header {
       flex-direction: column;
       align-items: flex-start;
       gap: 10px;
     }
     
     .issue-meta {
       grid-template-columns: 1fr;
     }
   }
   </style>

