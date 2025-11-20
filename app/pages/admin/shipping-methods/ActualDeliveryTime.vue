<template>
  <div class="actual-delivery-time">
       <div class="section-header">
         <h3>زمان تحویل واقعی</h3>
         <p>تحلیل عملکرد زمان تحویل و بهینه‌سازی</p>
       </div>
   
       <!-- Date Range Selector -->
       <div class="date-range-selector">
         <div class="form-group">
           <label>بازه زمانی:</label>
           <select v-model="selectedPeriod" @change="updateDeliveryData">
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
             <input v-model="customStartDate" type="date" @change="updateDeliveryData">
           </div>
           <div class="form-group">
             <label>تا تاریخ:</label>
             <input v-model="customEndDate" type="date" @change="updateDeliveryData">
           </div>
         </div>
       </div>
   
       <!-- Delivery Performance Summary -->
       <div class="delivery-summary">
         <div class="card average-card">
           <div class="card-icon">
             <i class="fas fa-clock"></i>
           </div>
           <div class="card-content">
             <h4>میانگین زمان تحویل</h4>
             <div class="number">{{ deliveryData.averageDeliveryTime }} روز</div>
             <div class="trend" :class="deliveryData.timeChange > 0 ? 'negative' : 'positive'">
               {{ deliveryData.timeChange > 0 ? '+' : '' }}{{ deliveryData.timeChange }}%
             </div>
           </div>
         </div>
   
         <div class="card ontime-card">
           <div class="card-icon">
             <i class="fas fa-check-circle"></i>
           </div>
           <div class="card-content">
             <h4>تحویل به موقع</h4>
             <div class="number">{{ deliveryData.onTimeDelivery }}%</div>
             <div class="trend" :class="deliveryData.onTimeChange > 0 ? 'positive' : 'negative'">
               {{ deliveryData.onTimeChange > 0 ? '+' : '' }}{{ deliveryData.onTimeChange }}%
             </div>
           </div>
         </div>
   
         <div class="card delay-card">
           <div class="card-icon">
             <i class="fas fa-exclamation-triangle"></i>
           </div>
           <div class="card-content">
             <h4>تاخیر در تحویل</h4>
             <div class="number">{{ deliveryData.delayedDeliveries }}%</div>
             <div class="trend" :class="deliveryData.delayChange > 0 ? 'negative' : 'positive'">
               {{ deliveryData.delayChange > 0 ? '+' : '' }}{{ deliveryData.delayChange }}%
             </div>
           </div>
         </div>
   
         <div class="card efficiency-card">
           <div class="card-icon">
             <i class="fas fa-tachometer-alt"></i>
           </div>
           <div class="card-content">
             <h4>کارایی تحویل</h4>
             <div class="number">{{ deliveryData.deliveryEfficiency }}%</div>
             <div class="trend" :class="deliveryData.efficiencyChange > 0 ? 'positive' : 'negative'">
               {{ deliveryData.efficiencyChange > 0 ? '+' : '' }}{{ deliveryData.efficiencyChange }}%
             </div>
           </div>
         </div>
       </div>
   
       <!-- Delivery Time Distribution -->
       <div class="time-distribution">
         <h4>توزیع زمان تحویل</h4>
         <div class="distribution-chart">
           <div class="time-bars">
             <div v-for="timeRange in timeDistribution" :key="timeRange.range" class="time-bar">
               <div class="time-label">
                 <span class="range">{{ timeRange.range }}</span>
                 <span class="count">({{ timeRange.count }})</span>
               </div>
               <div class="bar-container">
                 <div class="bar" :style="{ width: timeRange.percentage + '%', backgroundColor: timeRange.color }"></div>
                 <span class="percentage">{{ timeRange.percentage }}%</span>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Delivery Performance by Method -->
       <div class="method-performance">
         <h4>عملکرد تحویل بر اساس روش ارسال</h4>
         <div class="method-grid">
           <div v-for="method in methodPerformance" :key="method.id" class="method-card">
             <div class="method-header">
               <h5>{{ method.name }}</h5>
               <span class="method-type">{{ method.type }}</span>
             </div>
             <div class="method-stats">
               <div class="stat">
                 <span class="label">میانگین زمان:</span>
                 <span class="value">{{ method.averageTime }} روز</span>
               </div>
               <div class="stat">
                 <span class="label">تحویل به موقع:</span>
                 <span class="value">{{ method.onTimeRate }}%</span>
               </div>
               <div class="stat">
                 <span class="label">تاخیر:</span>
                 <span class="value" :class="method.delayRate > 10 ? 'negative' : 'positive'">
                   {{ method.delayRate }}%
                 </span>
               </div>
               <div class="stat">
                 <span class="label">تغییر:</span>
                 <span class="value" :class="method.timeChange >= 0 ? 'negative' : 'positive'">
                   {{ method.timeChange >= 0 ? '+' : '' }}{{ method.timeChange }}%
                 </span>
               </div>
             </div>
             <div class="performance-indicator">
               <div class="indicator-bar">
                 <div class="indicator" :style="{ width: method.performanceScore + '%', backgroundColor: getPerformanceColor(method.performanceScore) }"></div>
               </div>
               <span class="score">{{ method.performanceScore }}/100</span>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Delivery Time Trends -->
       <div class="delivery-trends">
         <h4>روند زمان تحویل</h4>
         <div class="trends-grid">
           <div class="trend-chart">
             <h5>روند روزانه تحویل</h5>
             <div class="chart-wrapper">
               <canvas ref="dailyTrendChart" width="400" height="200"></canvas>
             </div>
           </div>
   
           <div class="trend-chart">
             <h5>مقایسه روش‌های ارسال</h5>
             <div class="chart-wrapper">
               <canvas ref="methodComparisonChart" width="400" height="200"></canvas>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Delivery Issues Analysis -->
       <div class="delivery-issues">
         <h4>تحلیل مشکلات تحویل</h4>
         <div class="issues-grid">
           <div class="issue-card">
             <h5>دلایل تاخیر</h5>
             <div class="issue-list">
               <div v-for="issue in deliveryIssues" :key="issue.reason" class="issue-item">
                 <div class="issue-info">
                   <span class="issue-reason">{{ issue.reason }}</span>
                   <span class="issue-count">{{ issue.count }} مورد</span>
                 </div>
                 <div class="issue-bar">
                   <div class="bar" :style="{ width: issue.percentage + '%' }"></div>
                   <span class="percentage">{{ issue.percentage }}%</span>
                 </div>
               </div>
             </div>
           </div>
   
           <div class="issue-card">
             <h5>مناطق مشکل‌دار</h5>
             <div class="region-list">
               <div v-for="region in problemRegions" :key="region.name" class="region-item">
                 <div class="region-info">
                   <span class="region-name">{{ region.name }}</span>
                   <span class="region-delay">{{ region.averageDelay }} روز تاخیر</span>
                 </div>
                 <div class="region-bar">
                   <div class="bar" :style="{ width: region.delayPercentage + '%' }"></div>
                   <span class="percentage">{{ region.delayPercentage }}%</span>
                 </div>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Recent Deliveries -->
       <div class="recent-deliveries">
         <h4>آخرین تحویل‌ها</h4>
         <div class="deliveries-list">
           <div v-for="delivery in recentDeliveries" :key="delivery.id" class="delivery-item">
             <div class="delivery-header">
               <div class="delivery-info">
                 <span class="order-id">سفارش: {{ delivery.orderId }}</span>
                 <span class="customer-name">{{ delivery.customerName }}</span>
               </div>
               <div class="delivery-status" :class="delivery.status">
                 <i :class="getStatusIcon(delivery.status)"></i>
                 <span>{{ getStatusText(delivery.status) }}</span>
               </div>
             </div>
             <div class="delivery-details">
               <div class="detail">
                 <span class="label">روش ارسال:</span>
                 <span class="value">{{ delivery.shippingMethod }}</span>
               </div>
               <div class="detail">
                 <span class="label">زمان پیش‌بینی:</span>
                 <span class="value">{{ delivery.estimatedTime }} روز</span>
               </div>
               <div class="detail">
                 <span class="label">زمان واقعی:</span>
                 <span class="value" :class="delivery.actualTime > delivery.estimatedTime ? 'negative' : 'positive'">
                   {{ delivery.actualTime }} روز
                 </span>
               </div>
               <div class="detail">
                 <span class="label">تاریخ تحویل:</span>
                 <span class="value">{{ formatDate(delivery.deliveryDate) }}</span>
               </div>
             </div>
           </div>
         </div>
       </div>
   
       <!-- Export Options -->
       <div class="export-section">
         <h4>خروجی گزارش تحویل</h4>
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
   
   <script setup lang="ts">
   import { nextTick, onMounted, reactive, ref } from 'vue';
   
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
   
   const deliveryData = reactive({
     averageDeliveryTime: 0,
     timeChange: 0,
     onTimeDelivery: 0,
     onTimeChange: 0,
     delayedDeliveries: 0,
     delayChange: 0,
     deliveryEfficiency: 0,
     efficiencyChange: 0
   })
   
   const timeDistribution = ref([
     { range: '1 روز', count: 0, percentage: 0, color: '#27ae60' },
     { range: '2-3 روز', count: 0, percentage: 0, color: '#3498db' },
     { range: '4-5 روز', count: 0, percentage: 0, color: '#f39c12' },
     { range: '6-7 روز', count: 0, percentage: 0, color: '#e67e22' },
     { range: 'بیش از 7 روز', count: 0, percentage: 0, color: '#e74c3c' }
   ])
   
   const methodPerformance = ref([
     {
       id: 1,
       name: 'ارسال استاندارد',
       type: 'زمینی',
       averageTime: 3.2,
       onTimeRate: 85,
       delayRate: 15,
       timeChange: -0.5,
       performanceScore: 78
     },
     {
       id: 2,
       name: 'ارسال سریع',
       type: 'هوایی',
       averageTime: 1.1,
       onTimeRate: 95,
       delayRate: 5,
       timeChange: -0.2,
       performanceScore: 92
     },
     {
       id: 3,
       name: 'ارسال رایگان',
       type: 'زمینی',
       averageTime: 5.8,
       onTimeRate: 72,
       delayRate: 28,
       timeChange: 1.2,
       performanceScore: 65
     }
   ])
   
   const deliveryIssues = ref([
     { reason: 'مشکلات ترافیک', count: 45, percentage: 35 },
     { reason: 'تاخیر در بارگیری', count: 28, percentage: 22 },
     { reason: 'مشکلات آب و هوا', count: 20, percentage: 15 },
     { reason: 'مشکلات فنی', count: 18, percentage: 14 },
     { reason: 'سایر موارد', count: 19, percentage: 14 }
   ])
   
   const problemRegions = ref([
     { name: 'تهران', averageDelay: 2.5, delayPercentage: 25 },
     { name: 'اصفهان', averageDelay: 1.8, delayPercentage: 18 },
     { name: 'مشهد', averageDelay: 2.1, delayPercentage: 21 },
     { name: 'شیراز', averageDelay: 1.5, delayPercentage: 15 },
     { name: 'تبریز', averageDelay: 2.3, delayPercentage: 23 }
   ])
   
   const recentDeliveries = ref([
     {
       id: 1,
       orderId: 'ORD-001234',
       customerName: 'علی احمدی',
       status: 'delivered',
       shippingMethod: 'ارسال سریع',
       estimatedTime: 1,
       actualTime: 1,
       deliveryDate: new Date('2024-01-15')
     },
     {
       id: 2,
       orderId: 'ORD-001235',
       customerName: 'فاطمه محمدی',
       status: 'delayed',
       shippingMethod: 'ارسال استاندارد',
       estimatedTime: 3,
       actualTime: 5,
       deliveryDate: new Date('2024-01-14')
     },
     {
       id: 3,
       orderId: 'ORD-001236',
       customerName: 'محمد رضایی',
       status: 'delivered',
       shippingMethod: 'ارسال رایگان',
       estimatedTime: 5,
       actualTime: 6,
       deliveryDate: new Date('2024-01-13')
     }
   ])
   
   // Chart references
   const dailyTrendChart = ref<HTMLCanvasElement>()
   const methodComparisonChart = ref<HTMLCanvasElement>()
   
   // Methods
   const updateDeliveryData = async () => {
     try {
       // Simulate API call
       await new Promise(resolve => setTimeout(resolve, 1000))
       
       // Update delivery data
       deliveryData.averageDeliveryTime = Math.random() * 3 + 2 // 2 to 5 days
       deliveryData.timeChange = Math.floor(Math.random() * 20) - 10
       deliveryData.onTimeDelivery = Math.floor(Math.random() * 30) + 70
       deliveryData.onTimeChange = Math.floor(Math.random() * 15) - 7
       deliveryData.delayedDeliveries = 100 - deliveryData.onTimeDelivery
       deliveryData.delayChange = -deliveryData.onTimeChange
       deliveryData.deliveryEfficiency = Math.floor(Math.random() * 20) + 80
       deliveryData.efficiencyChange = Math.floor(Math.random() * 15) - 7
       
       // Update time distribution
       const total = 1000
       timeDistribution.value[0].count = Math.floor(total * 0.25)
       timeDistribution.value[0].percentage = 25
       timeDistribution.value[1].count = Math.floor(total * 0.35)
       timeDistribution.value[1].percentage = 35
       timeDistribution.value[2].count = Math.floor(total * 0.25)
       timeDistribution.value[2].percentage = 25
       timeDistribution.value[3].count = Math.floor(total * 0.10)
       timeDistribution.value[3].percentage = 10
       timeDistribution.value[4].count = Math.floor(total * 0.05)
       timeDistribution.value[4].percentage = 5
       
       updateCharts()
     } catch (error) {
       console.error('Error updating delivery data:', error)
     }
   }
   
   const updateCharts = () => {
     nextTick(() => {
       createDailyTrendChart()
       createMethodComparisonChart()
     })
   }
   
   const createDailyTrendChart = () => {
     if (!dailyTrendChart.value) return
     
     const ctx = dailyTrendChart.value.getContext('2d')
     if (!ctx) return
     
     // Simple line chart implementation
     const canvas = dailyTrendChart.value
     const width = canvas.width
     const height = canvas.height
     
     ctx.clearRect(0, 0, width, height)
     
     // Generate sample trend data
     const data = Array.from({ length: 30 }, () => Math.random() * 3 + 1) // 1 to 4 days
     const maxValue = Math.max(...data)
     
     ctx.strokeStyle = '#3498db'
     ctx.lineWidth = 2
     ctx.beginPath()
     
     data.forEach((value, index) => {
       const x = (index / (data.length - 1)) * (width - 40) + 20
       const y = height - ((value / maxValue) * (height - 40) + 20)
       
       if (index === 0) {
         ctx.moveTo(x, y)
       } else {
         ctx.lineTo(x, y)
       }
     })
     
     ctx.stroke()
   }
   
   const createMethodComparisonChart = () => {
     if (!methodComparisonChart.value) return
     
     const ctx = methodComparisonChart.value.getContext('2d')
     if (!ctx) return
     
     // Simple bar chart implementation
     const canvas = methodComparisonChart.value
     const width = canvas.width
     const height = canvas.height
     
     ctx.clearRect(0, 0, width, height)
     
     const data = methodPerformance.value.map(item => item.averageTime)
     const maxValue = Math.max(...data)
     const barWidth = width / data.length - 20
     
     data.forEach((value, index) => {
       const barHeight = (value / maxValue) * (height - 40)
       const x = index * (barWidth + 20) + 10
       const y = height - barHeight - 20
       
       ctx.fillStyle = ['#27ae60', '#3498db', '#f39c12'][index]
       ctx.fillRect(x, y, barWidth, barHeight)
       
       // Draw value label
       ctx.fillStyle = '#2c3e50'
       ctx.font = '12px Arial'
       ctx.textAlign = 'center'
       ctx.fillText(value.toFixed(1) + ' روز', x + barWidth / 2, y - 5)
     })
   }
   
   const getPerformanceColor = (score: number): string => {
     if (score >= 90) return '#27ae60'
     if (score >= 80) return '#3498db'
     if (score >= 70) return '#f39c12'
     return '#e74c3c'
   }
   
   const getStatusIcon = (status: string): string => {
     switch (status) {
       case 'delivered': return 'fas fa-check-circle'
       case 'delayed': return 'fas fa-exclamation-triangle'
       case 'in_transit': return 'fas fa-truck'
       default: return 'fas fa-clock'
     }
   }
   
   const getStatusText = (status: string): string => {
     switch (status) {
       case 'delivered': return 'تحویل شده'
       case 'delayed': return 'تاخیر'
       case 'in_transit': return 'در حال ارسال'
       default: return 'در انتظار'
     }
   }
   
   const formatDate = (date: Date): string => {
     return new Intl.DateTimeFormat('fa-IR').format(date)
   }
   
   const exportToExcel = () => {
     const data = {
       deliveryData,
       timeDistribution: timeDistribution.value,
       methodPerformance: methodPerformance.value,
       deliveryIssues: deliveryIssues.value,
       problemRegions: problemRegions.value,
       recentDeliveries: recentDeliveries.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'excel')
   }
   
   const exportToPDF = () => {
     const data = {
       deliveryData,
       timeDistribution: timeDistribution.value,
       methodPerformance: methodPerformance.value,
       deliveryIssues: deliveryIssues.value,
       problemRegions: problemRegions.value,
       recentDeliveries: recentDeliveries.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'pdf')
   }
   
   const printReport = () => {
     window.print()
   }
   
   // Lifecycle
   onMounted(() => {
     updateDeliveryData()
   })
   </script>
   
   <style scoped>
   .actual-delivery-time {
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
   
   .delivery-summary {
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
   
   .average-card {
     background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
   }
   
   .ontime-card {
     background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
   }
   
   .delay-card {
     background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
   }
   
   .efficiency-card {
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
   
   .time-distribution {
     margin-bottom: 30px;
   }
   
   .time-distribution h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .distribution-chart {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .time-bars {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .time-bar {
     display: flex;
     align-items: center;
     gap: 15px;
   }
   
   .time-label {
     display: flex;
     align-items: center;
     gap: 8px;
     min-width: 120px;
   }
   
   .range {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .count {
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
     border-radius: 10px;
     min-width: 50px;
   }
   
   .percentage {
     font-size: 12px;
     color: #2c3e50;
     min-width: 40px;
   }
   
   .method-performance {
     margin-bottom: 30px;
   }
   
   .method-performance h4 {
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
   
   .method-stats {
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
   
   .stat .value.positive {
     color: #27ae60;
   }
   
   .stat .value.negative {
     color: #e74c3c;
   }
   
   .performance-indicator {
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .indicator-bar {
     flex: 1;
     height: 8px;
     background: #e9ecef;
     border-radius: 4px;
     overflow: hidden;
   }
   
   .indicator {
     height: 100%;
     border-radius: 4px;
     transition: width 0.3s ease;
   }
   
   .score {
     font-size: 12px;
     color: #6c757d;
     min-width: 50px;
   }
   
   .delivery-trends {
     margin-bottom: 30px;
   }
   
   .delivery-trends h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .trends-grid {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
     gap: 30px;
   }
   
   .trend-chart {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .trend-chart h5 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .chart-wrapper {
     position: relative;
     height: 300px;
   }
   
   .delivery-issues {
     margin-bottom: 30px;
   }
   
   .delivery-issues h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .issues-grid {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
     gap: 30px;
   }
   
   .issue-card {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .issue-card h5 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .issue-list,
   .region-list {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .issue-item,
   .region-item {
     display: flex;
     flex-direction: column;
     gap: 8px;
   }
   
   .issue-info,
   .region-info {
     display: flex;
     justify-content: space-between;
     align-items: center;
     font-size: 14px;
   }
   
   .issue-reason,
   .region-name {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .issue-count,
   .region-delay {
     color: #6c757d;
     font-size: 12px;
   }
   
   .issue-bar,
   .region-bar {
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .issue-bar .bar,
   .region-bar .bar {
     height: 12px;
     background: #e74c3c;
     border-radius: 6px;
     min-width: 100px;
   }
   
   .recent-deliveries {
     margin-bottom: 30px;
   }
   
   .recent-deliveries h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .deliveries-list {
     display: flex;
     flex-direction: column;
     gap: 15px;
   }
   
   .delivery-item {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
     border: 1px solid #e9ecef;
   }
   
   .delivery-header {
     display: flex;
     justify-content: space-between;
     align-items: center;
     margin-bottom: 15px;
   }
   
   .delivery-info {
     display: flex;
     flex-direction: column;
     gap: 5px;
   }
   
   .order-id {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .customer-name {
     font-size: 12px;
     color: #6c757d;
   }
   
   .delivery-status {
     display: flex;
     align-items: center;
     gap: 8px;
     padding: 6px 12px;
     border-radius: 20px;
     font-size: 12px;
     font-weight: 600;
   }
   
   .delivery-status.delivered {
     background: rgba(39, 174, 96, 0.1);
     color: #27ae60;
   }
   
   .delivery-status.delayed {
     background: rgba(231, 76, 60, 0.1);
     color: #e74c3c;
   }
   
   .delivery-status.in_transit {
     background: rgba(52, 152, 219, 0.1);
     color: #3498db;
   }
   
   .delivery-details {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
     gap: 10px;
   }
   
   .detail {
     display: flex;
     justify-content: space-between;
     font-size: 14px;
   }
   
   .detail .label {
     color: #6c757d;
   }
   
   .detail .value {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .detail .value.positive {
     color: #27ae60;
   }
   
   .detail .value.negative {
     color: #e74c3c;
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
     .delivery-summary {
       grid-template-columns: 1fr;
     }
     
     .method-grid {
       grid-template-columns: 1fr;
     }
     
     .trends-grid {
       grid-template-columns: 1fr;
     }
     
     .issues-grid {
       grid-template-columns: 1fr;
     }
     
     .custom-date-range {
       flex-direction: column;
     }
     
     .export-buttons {
       flex-direction: column;
     }
     
     .delivery-header {
       flex-direction: column;
       align-items: flex-start;
       gap: 10px;
     }
     
     .delivery-details {
       grid-template-columns: 1fr;
     }
   }
   </style> 
