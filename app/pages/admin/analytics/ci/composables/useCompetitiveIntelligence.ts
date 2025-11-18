import { ref, computed } from 'vue'

// Types
export interface MarketKPI {
     marketShare: number
     marketShareChange: number
     activeCompetitors: number
     newCompetitors: number
     marketGrowth: number
     marketGrowthChange: number
     threatLevel: string
     threatDescription: string
}

export interface MarketStatus {
     id: number
     title: string
     level: 'stable' | 'high' | 'medium' | 'low'
     value: string
}

export interface KeyTrend {
     id: number
     title: string
     description: string
     direction: 'up' | 'down'
     time: string
}

export interface Alert {
     id: number
     title: string
     description: string
     priority: 'high' | 'medium' | 'low'
     competitor: string
     timestamp: Date
}

export interface CompetitorProfile {
     id: number
     name: string
     website: string
     foundedYear: number
     companySize: string
     targetMarkets: string[]
     keyProducts: string[]
     pricing: string
     strengths: string[]
     weaknesses: string[]
     recentChanges: string[]
     threatLevel: 'low' | 'medium' | 'high'
}

// Composable
export const useCompetitiveIntelligence = () => {
     // State
     const isLoading = ref(false)
     const error = ref<string | null>(null)
     const lastUpdateTime = ref<Date>(new Date())

     // Data
     const marketKPI = ref<MarketKPI>({
          marketShare: 23.5,
          marketShareChange: 2.1,
          activeCompetitors: 8,
          newCompetitors: 1,
          marketGrowth: 12.8,
          marketGrowthChange: -0.5,
          threatLevel: 'متوسط',
          threatDescription: 'نیاز به توجه بیشتر'
     })

     const marketStatuses = ref<MarketStatus[]>([
          { id: 1, title: 'وضعیت کلی بازار', level: 'stable', value: 'پایدار' },
          { id: 2, title: 'رقابت', level: 'high', value: 'شدید' },
          { id: 3, title: 'نوآوری', level: 'medium', value: 'متوسط' },
          { id: 4, title: 'ثبات قیمت', level: 'stable', value: 'پایدار' }
     ])

     const keyTrends = ref<KeyTrend[]>([
          { id: 1, title: 'افزایش استفاده از هوش مصنوعی', description: 'رقبا در حال سرمایه‌گذاری روی AI هستند', direction: 'up', time: '2 ساعت پیش' },
          { id: 2, title: 'کاهش قیمت محصولات', description: 'رقبای اصلی قیمت‌ها را کاهش داده‌اند', direction: 'down', time: '1 روز پیش' },
          { id: 3, title: 'ورود به بازارهای جدید', description: 'رقبا در حال گسترش جغرافیایی هستند', direction: 'up', time: '3 روز پیش' }
     ])

     const alerts = ref<Alert[]>([
          { id: 1, title: 'محصول جدید رقیب', description: 'شرکت ABC محصول جدیدی در حوزه IoT معرفی کرد', priority: 'high', competitor: 'شرکت ABC', timestamp: new Date() },
          { id: 2, title: 'کمپین تبلیغاتی', description: 'رقبای اصلی کمپین تبلیغاتی جدیدی راه‌اندازی کرده‌اند', priority: 'medium', competitor: 'رقبای اصلی', timestamp: new Date(Date.now() - 3600000) },
          { id: 3, title: 'تغییر قیمت', description: 'شرکت XYZ قیمت محصولات خود را 15% کاهش داد', priority: 'high', competitor: 'شرکت XYZ', timestamp: new Date(Date.now() - 7200000) },
          { id: 4, title: 'شراکت جدید', description: 'رقبای اصلی با شرکت‌های فناوری شراکت کرده‌اند', priority: 'low', competitor: 'رقبای اصلی', timestamp: new Date(Date.now() - 86400000) },
          { id: 5, title: 'تغییر استراتژی', description: 'شرکت DEF استراتژی بازاریابی خود را تغییر داده است', priority: 'medium', competitor: 'شرکت DEF', timestamp: new Date(Date.now() - 172800000) }
     ])

     const competitors = ref<CompetitorProfile[]>([
          {
               id: 1,
               name: 'شرکت ABC',
               website: 'https://abc.com',
               foundedYear: 2015,
               companySize: '100-500 نفر',
               targetMarkets: ['ایران', 'ترکیه', 'امارات'],
               keyProducts: ['IoT Platform', 'Smart Home', 'Industrial Sensors'],
               pricing: 'متوسط تا بالا',
               strengths: ['تکنولوژی پیشرفته', 'پشتیبانی قوی', 'کیفیت بالا'],
               weaknesses: ['قیمت بالا', 'زمان تحویل طولانی'],
               recentChanges: ['لانچ محصول جدید IoT', 'ورود به بازار ترکیه'],
               threatLevel: 'high'
          },
          {
               id: 2,
               name: 'شرکت XYZ',
               website: 'https://xyz.com',
               foundedYear: 2010,
               companySize: '500-1000 نفر',
               targetMarkets: ['ایران', 'عراق', 'افغانستان'],
               keyProducts: ['Mobile Apps', 'E-commerce', 'Cloud Services'],
               pricing: 'متوسط',
               strengths: ['تجربه طولانی', 'شبکه توزیع قوی', 'قیمت رقابتی'],
               weaknesses: ['نوآوری کم', 'خدمات پس از فروش ضعیف'],
               recentChanges: ['کاهش 15% قیمت‌ها', 'کمپین تبلیغاتی جدید'],
               threatLevel: 'medium'
          }
     ])

     // Computed
     const recentAlerts = computed(() => {
          return alerts.value.slice(0, 3)
     })

     const highPriorityAlerts = computed(() => {
          return alerts.value.filter(alert => alert.priority === 'high')
     })

     const threatLevelColor = computed(() => {
          switch (marketKPI.value.threatLevel) {
               case 'کم': return 'text-green-600'
               case 'متوسط': return 'text-yellow-600'
               case 'زیاد': return 'text-red-600'
               default: return 'text-gray-600'
          }
     })

     const threatLevelBgColor = computed(() => {
          switch (marketKPI.value.threatLevel) {
               case 'کم': return 'bg-gradient-to-br from-green-500 to-green-600'
               case 'متوسط': return 'bg-gradient-to-br from-yellow-500 to-yellow-600'
               case 'زیاد': return 'bg-gradient-to-br from-red-500 to-red-600'
               default: return 'bg-gradient-to-br from-gray-500 to-gray-600'
          }
     })

     // Methods
     const refreshData = async () => {
          isLoading.value = true
          error.value = null

          try {
               // TODO: در پروژه واقعی، اینجا API call انجام می‌شود
               // Update last update time
               lastUpdateTime.value = new Date()

               // TODO: Update data from API response
               console.log('CI data refreshed successfully')
          } catch (err) {
               error.value = 'خطا در به‌روزرسانی داده‌ها'
               console.error('Error refreshing CI data:', err)
          } finally {
               isLoading.value = false
          }
     }

     const getCompetitorById = (id: number) => {
          return competitors.value.find(comp => comp.id === id)
     }

     const getAlertsByPriority = (priority: 'high' | 'medium' | 'low') => {
          return alerts.value.filter(alert => alert.priority === priority)
     }

     const getAlertsByCompetitor = (competitorName: string) => {
          return alerts.value.filter(alert => alert.competitor === competitorName)
     }

     const addAlert = (alert: Omit<Alert, 'id' | 'timestamp'>) => {
          const newAlert: Alert = {
               ...alert,
               id: Date.now(),
               timestamp: new Date()
          }
          alerts.value.unshift(newAlert)
     }

     const updateMarketKPI = (updates: Partial<MarketKPI>) => {
          marketKPI.value = { ...marketKPI.value, ...updates }
     }

     const addCompetitor = (competitor: Omit<CompetitorProfile, 'id'>) => {
          const newCompetitor: CompetitorProfile = {
               ...competitor,
               id: Date.now()
          }
          competitors.value.push(newCompetitor)
     }

     const updateCompetitor = (id: number, updates: Partial<CompetitorProfile>) => {
          const index = competitors.value.findIndex(comp => comp.id === id)
          if (index !== -1) {
               competitors.value[index] = { ...competitors.value[index], ...updates }
          }
     }

     const deleteCompetitor = (id: number) => {
          const index = competitors.value.findIndex(comp => comp.id === id)
          if (index !== -1) {
               competitors.value.splice(index, 1)
          }
     }

     // Utility functions
     const formatTime = (date: Date) => {
          const now = new Date()
          const diff = now.getTime() - date.getTime()
          const minutes = Math.floor(diff / 60000)

          if (minutes < 1) return 'همین الان'
          if (minutes < 60) return `${minutes} دقیقه پیش`
          if (minutes < 1440) return `${Math.floor(minutes / 60)} ساعت پیش`
          return `${Math.floor(minutes / 1440)} روز پیش`
     }

     const getPriorityColor = (priority: 'high' | 'medium' | 'low') => {
          switch (priority) {
               case 'high': return 'bg-red-500'
               case 'medium': return 'bg-yellow-500'
               case 'low': return 'bg-blue-500'
               default: return 'bg-gray-500'
          }
     }

     const getPriorityBadgeClass = (priority: 'high' | 'medium' | 'low') => {
          switch (priority) {
               case 'high': return 'bg-red-100 text-red-800'
               case 'medium': return 'bg-yellow-100 text-yellow-800'
               case 'low': return 'bg-blue-100 text-blue-800'
               default: return 'bg-gray-100 text-gray-800'
          }
     }

     const getPriorityText = (priority: 'high' | 'medium' | 'low') => {
          switch (priority) {
               case 'high': return 'زیاد'
               case 'medium': return 'متوسط'
               case 'low': return 'کم'
               default: return 'نامشخص'
          }
     }

     const getStatusColor = (level: string) => {
          switch (level) {
               case 'stable': return 'bg-green-500'
               case 'high': return 'bg-red-500'
               case 'medium': return 'bg-yellow-500'
               default: return 'bg-gray-500'
          }
     }

     return {
          // State
          isLoading,
          error,
          lastUpdateTime,

          // Data
          marketKPI,
          marketStatuses,
          keyTrends,
          alerts,
          competitors,

          // Computed
          recentAlerts,
          highPriorityAlerts,
          threatLevelColor,
          threatLevelBgColor,

          // Methods
          refreshData,
          getCompetitorById,
          getAlertsByPriority,
          getAlertsByCompetitor,
          addAlert,
          updateMarketKPI,
          addCompetitor,
          updateCompetitor,
          deleteCompetitor,

          // Utility functions
          formatTime,
          getPriorityColor,
          getPriorityBadgeClass,
          getPriorityText,
          getStatusColor
     }
}
