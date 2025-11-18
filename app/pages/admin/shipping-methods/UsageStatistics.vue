<template>
  <div class="usage-statistics">
       <div class="section-header">
         <h3>آمار استفاده</h3>
         <p>نمایش آمار استفاده از روش‌های ارسال</p>
       </div>
   
       <!-- Date Range Selector -->
       <div class="date-range-selector">
         <div class="form-group">
           <label>بازه زمانی:</label>
           <select v-model="selectedPeriod" @change="updateStatistics">
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
             <input type="date" v-model="customStartDate" @change="updateStatistics">
           </div>
           <div class="form-group">
             <label>تا تاریخ:</label>
             <input type="date" v-model="customEndDate" @change="updateStatistics">
           </div>
         </div>
       </div>
   
       <!-- Summary Cards -->
       <div class="summary-cards">
         <div class="card">
           <div class="card-icon">
             <i class="fas fa-shipping-fast"></i>
           </div>
           <div class="card-content">
             <h4>کل سفارشات</h4>
             <div class="number">{{ formatNumber(statistics.totalOrders) }}</div>
             <div class="trend" :class="statistics.ordersTrend > 0 ? 'positive' : 'negative'">
               {{ statistics.ordersTrend > 0 ? '+' : '' }}{{ statistics.ordersTrend }}%
             </div>
           </div>
         </div>
   
         <div class="card">
           <div class="card-icon">
             <i class="fas fa-dollar-sign"></i>
           </div>
           <div class="card-content">
             <h4>درآمد کل</h4>
             <div class="number">{{ formatCurrency(statistics.totalRevenue) }}</div>
             <div class="trend" :class="statistics.revenueTrend > 0 ? 'positive' : 'negative'">
               {{ statistics.revenueTrend > 0 ? '+' : '' }}{{ statistics.revenueTrend }}%
             </div>
           </div>
         </div>
   
         <div class="card">
           <div class="card-icon">
             <i class="fas fa-users"></i>
           </div>
           <div class="card-content">
             <h4>مشتریان فعال</h4>
             <div class="number">{{ formatNumber(statistics.activeCustomers) }}</div>
             <div class="trend" :class="statistics.customersTrend > 0 ? 'positive' : 'negative'">
               {{ statistics.customersTrend > 0 ? '+' : '' }}{{ statistics.customersTrend }}%
             </div>
           </div>
         </div>
   
         <div class="card">
           <div class="card-icon">
             <i class="fas fa-chart-line"></i>
           </div>
           <div class="card-content">
             <h4>نرخ رشد</h4>
             <div class="number">{{ statistics.growthRate }}%</div>
             <div class="trend" :class="statistics.growthRate > 0 ? 'positive' : 'negative'">
               نسبت به دوره قبل
             </div>
           </div>
         </div>
       </div>
   
       <!-- Charts Section -->
       <div class="charts-section">
         <!-- Usage by Shipping Method -->
         <div class="chart-container">
           <h4>استفاده بر اساس روش ارسال</h4>
           <div class="chart-wrapper">
             <canvas ref="usageChart" width="400" height="200"></canvas>
           </div>
         </div>
   
         <!-- Daily Usage Trend -->
         <div class="chart-container">
           <h4>روند روزانه استفاده</h4>
           <div class="chart-wrapper">
             <canvas ref="trendChart" width="400" height="200"></canvas>
           </div>
         </div>
       </div>
   
       <!-- Detailed Statistics Table -->
       <div class="detailed-stats">
         <h4>آمار تفصیلی</h4>
         <div class="table-container">
           <table>
             <thead>
               <tr>
                 <th>روش ارسال</th>
                 <th>تعداد سفارش</th>
                 <th>درصد استفاده</th>
                 <th>میانگین ارزش</th>
                 <th>میانگین زمان تحویل</th>
                 <th>نرخ رضایت</th>
               </tr>
             </thead>
             <tbody>
               <tr v-for="method in detailedStats" :key="method.id">
                 <td>
                   <div class="method-info">
                     <span class="method-name">{{ method.name }}</span>
                     <span class="method-type">{{ method.type }}</span>
                   </div>
                 </td>
                 <td>{{ formatNumber(method.orderCount) }}</td>
                 <td>
                   <div class="percentage-bar">
                     <div class="bar" :style="{ width: method.usagePercentage + '%' }"></div>
                     <span>{{ method.usagePercentage }}%</span>
                   </div>
                 </td>
                 <td>{{ formatCurrency(method.averageValue) }}</td>
                 <td>{{ method.averageDeliveryTime }} روز</td>
                 <td>
                   <div class="satisfaction-rating">
                     <div class="stars">
                       <i v-for="star in 5" :key="star" 
                          :class="star <= method.satisfactionRating ? 'fas fa-star filled' : 'far fa-star'"></i>
                     </div>
                     <span>{{ method.satisfactionRating }}/5</span>
                   </div>
                 </td>
               </tr>
             </tbody>
           </table>
         </div>
       </div>
   
       <!-- Export Options -->
       <div class="export-section">
         <h4>خروجی گزارش</h4>
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
   import { ref, reactive, onMounted, nextTick } from 'vue'
   import Chart from 'chart.js/auto'
   
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
   
   const statistics = reactive({
     totalOrders: 0,
     ordersTrend: 0,
     totalRevenue: 0,
     revenueTrend: 0,
     activeCustomers: 0,
     customersTrend: 0,
     growthRate: 0
   })
   
   const detailedStats = ref([
     {
       id: 1,
       name: 'ارسال استاندارد',
       type: 'زمینی',
       orderCount: 1250,
       usagePercentage: 45,
       averageValue: 850000,
       averageDeliveryTime: 3,
       satisfactionRating: 4.2
     },
     {
       id: 2,
       name: 'ارسال سریع',
       type: 'هوایی',
       orderCount: 890,
       usagePercentage: 32,
       averageValue: 1200000,
       averageDeliveryTime: 1,
       satisfactionRating: 4.5
     },
     {
       id: 3,
       name: 'ارسال رایگان',
       type: 'زمینی',
       orderCount: 620,
       usagePercentage: 23,
       averageValue: 650000,
       averageDeliveryTime: 5,
       satisfactionRating: 4.0
     }
   ])
   
   // Chart references
   const usageChart = ref<HTMLCanvasElement>()
   const trendChart = ref<HTMLCanvasElement>()
   let usageChartInstance: Chart | null = null
   let trendChartInstance: Chart | null = null
   
   // Methods
   const updateStatistics = async () => {
     try {
       // Simulate API call
       await new Promise(resolve => setTimeout(resolve, 1000))
       
       // Update statistics based on selected period
       statistics.totalOrders = Math.floor(Math.random() * 5000) + 1000
       statistics.ordersTrend = Math.floor(Math.random() * 20) - 10
       statistics.totalRevenue = Math.floor(Math.random() * 1000000000) + 500000000
       statistics.revenueTrend = Math.floor(Math.random() * 20) - 10
       statistics.activeCustomers = Math.floor(Math.random() * 1000) + 500
       statistics.customersTrend = Math.floor(Math.random() * 20) - 10
       statistics.growthRate = Math.floor(Math.random() * 30) - 5
       
       updateCharts()
     } catch (error) {
       console.error('Error updating statistics:', error)
     }
   }
   
   const updateCharts = () => {
     if (usageChartInstance) {
       usageChartInstance.destroy()
     }
     if (trendChartInstance) {
       trendChartInstance.destroy()
     }
     
     nextTick(() => {
       createUsageChart()
       createTrendChart()
     })
   }
   
   const createUsageChart = () => {
     if (!usageChart.value) return
     
     const ctx = usageChart.value.getContext('2d')
     if (!ctx) return
     
     usageChartInstance = new Chart(ctx, {
       type: 'doughnut',
       data: {
         labels: detailedStats.value.map(stat => stat.name),
         datasets: [{
           data: detailedStats.value.map(stat => stat.orderCount),
           backgroundColor: [
             '#FF6384',
             '#36A2EB',
             '#FFCE56',
             '#4BC0C0',
             '#9966FF'
           ]
         }]
       },
       options: {
         responsive: true,
         plugins: {
           legend: {
             position: 'bottom'
           }
         }
       }
     })
   }
   
   const createTrendChart = () => {
     if (!trendChart.value) return
     
     const ctx = trendChart.value.getContext('2d')
     if (!ctx) return
     
     const labels = Array.from({ length: 30 }, (_, i) => `روز ${i + 1}`)
     const data = Array.from({ length: 30 }, () => Math.floor(Math.random() * 100) + 50)
     
     trendChartInstance = new Chart(ctx, {
       type: 'line',
       data: {
         labels,
         datasets: [{
           label: 'تعداد سفارشات',
           data,
           borderColor: '#36A2EB',
           backgroundColor: 'rgba(54, 162, 235, 0.1)',
           tension: 0.4
         }]
       },
       options: {
         responsive: true,
         plugins: {
           legend: {
             display: false
           }
         },
         scales: {
           y: {
             beginAtZero: true
           }
         }
       }
     })
   }
   
   const formatNumber = (num: number): string => {
     return new Intl.NumberFormat('fa-IR').format(num)
   }
   
   const formatCurrency = (amount: number): string => {
     return new Intl.NumberFormat('fa-IR', {
       style: 'currency',
       currency: 'IRR'
     }).format(amount)
   }
   
   const exportToExcel = () => {
     const data = {
       statistics,
       detailedStats: detailedStats.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'excel')
   }
   
   const exportToPDF = () => {
     const data = {
       statistics,
       detailedStats: detailedStats.value,
       period: selectedPeriod.value
     }
     emit('exportData', data, 'pdf')
   }
   
   const printReport = () => {
     window.print()
   }
   
   // Lifecycle
   onMounted(() => {
     updateStatistics()
   })
   </script>
   
   <style scoped>
   .usage-statistics {
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
   
   .summary-cards {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
     gap: 20px;
     margin-bottom: 30px;
   }
   
   .card {
     background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
     color: white;
     padding: 20px;
     border-radius: 12px;
     display: flex;
     align-items: center;
     gap: 15px;
     box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
   }
   
   .card:nth-child(2) {
     background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
   }
   
   .card:nth-child(3) {
     background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
   }
   
   .card:nth-child(4) {
     background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
   }
   
   .card-icon {
     font-size: 24px;
     opacity: 0.8;
   }
   
   .card-content h4 {
     margin: 0 0 5px 0;
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
   
   .charts-section {
     display: grid;
     grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
     gap: 30px;
     margin-bottom: 30px;
   }
   
   .chart-container {
     background: #f8f9fa;
     padding: 20px;
     border-radius: 8px;
   }
   
   .chart-container h4 {
     margin-bottom: 20px;
     color: #2c3e50;
     text-align: center;
   }
   
   .chart-wrapper {
     position: relative;
     height: 300px;
   }
   
   .detailed-stats {
     margin-bottom: 30px;
   }
   
   .detailed-stats h4 {
     margin-bottom: 20px;
     color: #2c3e50;
   }
   
   .table-container {
     overflow-x: auto;
   }
   
   table {
     width: 100%;
     border-collapse: collapse;
     background: white;
     border-radius: 8px;
     overflow: hidden;
     box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
   }
   
   th, td {
     padding: 12px;
     text-align: right;
     border-bottom: 1px solid #eee;
   }
   
   th {
     background: #f8f9fa;
     font-weight: 600;
     color: #2c3e50;
   }
   
   .method-info {
     display: flex;
     flex-direction: column;
     gap: 2px;
   }
   
   .method-name {
     font-weight: 600;
     color: #2c3e50;
   }
   
   .method-type {
     font-size: 12px;
     color: #7f8c8d;
     background: #ecf0f1;
     padding: 2px 6px;
     border-radius: 4px;
     display: inline-block;
   }
   
   .percentage-bar {
     display: flex;
     align-items: center;
     gap: 10px;
   }
   
   .bar {
     height: 8px;
     background: #3498db;
     border-radius: 4px;
     min-width: 50px;
   }
   
   .satisfaction-rating {
     display: flex;
     align-items: center;
     gap: 8px;
   }
   
   .stars {
     display: flex;
     gap: 2px;
   }
   
   .stars i {
     font-size: 12px;
   }
   
   .stars i.filled {
     color: #f39c12;
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
     .summary-cards {
       grid-template-columns: 1fr;
     }
     
     .charts-section {
       grid-template-columns: 1fr;
     }
     
     .custom-date-range {
       flex-direction: column;
     }
     
     .export-buttons {
       flex-direction: column;
     }
   }
   </style>
