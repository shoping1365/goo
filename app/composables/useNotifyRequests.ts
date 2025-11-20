import { useRuntimeConfig } from 'nuxt/app'
import { computed, ref } from 'vue'

// تعریف useFetch برای Nuxt 3
// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare const useFetch: any

interface User {
  id: number
  name: string
  email: string
}

interface Product {
  id: number
  name: string
  code: string
}

type RequestType = 'stock' | 'discount' | 'price'
type RequestStatus = 'pending' | 'completed' | 'cancelled'

interface NotificationRequest {
  id: number
  user: User
  product: Product
  type: RequestType
  status: RequestStatus
  created_at: string
}

interface Settings {
  emailNotifications: boolean
  smsNotifications: boolean
  pushNotifications: boolean
  messageTemplate: string
  delayMinutes: number
  dailyLimit: number
  // تنظیمات ظاهر
  theme: 'light' | 'dark' | 'auto'
  primaryColor: string
  fontSize: 'small' | 'medium' | 'large'
  // تنظیمات ایمیل
  emailProvider: 'smtp' | 'mailgun' | 'sendgrid'
  smtpHost: string
  smtpPort: number
  smtpUsername: string
  smtpPassword: string
  emailFrom: string
  emailFromName: string
  // تنظیمات پیامک
  smsProvider: 'kavenegar' | 'melipayamak' | 'farapayamak'
  smsApiKey: string
  smsLineNumber: string
  // اطلاع‌رسانی خودکار
  autoNotifyStock: boolean
  autoNotifyDiscount: boolean
  autoNotifyDelay: number
  maxAutoNotifications: number
  // پنل کاربری
  userCanCancel: boolean
  userCanEdit: boolean
  showUserHistory: boolean
  userDashboardEnabled: boolean
}

interface StockRequest {
  id: number
  date: string
  time: string
  price: number
  phone: string
  product: string
  status: string
  notificationDate?: string
  notificationTime?: string
}

interface ApiNotification {
  id: number
  created_at: string
  phone?: string
  product?: { name: string }
  product_id: number
  status: string
}

export const useNotifyRequests = () => {
  // Reactive data
  const totalRequests = ref(847)
  const pendingRequests = ref(312)
  const completedRequests = ref(495)
  const todayRequests = ref(28)
  const successRate = ref(87.5)

  const activeTab = ref('requests')
  const activeNavItem = ref('dashboard')
  const searchQuery = ref('')
  const statusFilter = ref('')
  const typeFilter = ref('')
  const selectedRequests = ref<number[]>([])
  const showSettings = ref(false)
  const activeSettingsTab = ref('general')

  const settings = ref<Settings>({
    emailNotifications: true,
    smsNotifications: true,
    pushNotifications: false,
    messageTemplate: 'سلام {name}، محصول {product} که درخواست اطلاع‌رسانی کرده‌اید اکنون موجود است.',
    delayMinutes: 0,
    dailyLimit: 1000,
    // تنظیمات ظاهر
    theme: 'light',
    primaryColor: '#3B82F6',
    fontSize: 'medium',
    // تنظیمات ایمیل
    emailProvider: 'smtp',
    smtpHost: '',
    smtpPort: 587,
    smtpUsername: '',
    smtpPassword: '',
    emailFrom: 'noreply@example.com',
    emailFromName: 'فروشگاه ما',
    // تنظیمات پیامک
    smsProvider: 'kavenegar',
    smsApiKey: '',
    smsLineNumber: '',
    // اطلاع‌رسانی خودکار
    autoNotifyStock: true,
    autoNotifyDiscount: true,
    autoNotifyDelay: 5,
    maxAutoNotifications: 100,
    // پنل کاربری
    userCanCancel: true,
    userCanEdit: false,
    showUserHistory: true,
    userDashboardEnabled: true
  })

  const tabs = [
    { id: 'requests', name: 'درخواست‌ها' },
    { id: 'analytics', name: 'تحلیل و گزارش' },
    { id: 'settings', name: 'تنظیمات' }
  ]

  const settingsTabs = [
    { id: 'general', name: 'عمومی', icon: 'M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4' },
    { id: 'appearance', name: 'ظاهر', icon: 'M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zM21 5a2 2 0 00-2-2h-4a2 2 0 00-2 2v12a4 4 0 004 4h4a2 2 0 002-2V5z' },
    { id: 'content', name: 'متون', icon: 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z' },
    { id: 'email', name: 'ایمیل', icon: 'M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z' },
    { id: 'sms', name: 'پیامک', icon: 'M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z' },
    { id: 'automation', name: 'اطلاع‌رسانی خودکار', icon: 'M13 10V3L4 14h7v7l9-11h-7z' },
    { id: 'logs', name: 'لاگ‌ها', icon: 'M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01' },
    { id: 'user-panel', name: 'پنل کاربری', icon: 'M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z' }
  ]

  const navigationItems = [
    {
      id: 'dashboard',
      name: 'داشبورد',
      icon: 'M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2H5a2 2 0 00-2-2z M9 9h6v6H9z',
      count: null as number | null
    },
    {
      id: 'stock',
      name: 'موجودی',
      icon: 'M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10',
      count: null as number | null
    },
    {
      id: 'discount',
      name: 'تخفیف',
      icon: 'M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z',
      count: null as number | null
    },
    {
      id: 'stock-completed',
      name: 'اطلاع‌رسانی شده موجودی',
      icon: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z',
      count: null as number | null
    },
    {
      id: 'discount-completed',
      name: 'اطلاع‌رسانی شده تخفیف',
      icon: 'M5 13l4 4L19 7',
      count: null as number | null
    },
    {
      id: 'settings',
      name: 'تنظیمات',
      icon: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z M15 12a3 3 0 11-6 0 3 3 0 016 0z',
      count: null as number | null
    }
  ]

  const requests = ref<NotificationRequest[]>([
    {
      id: 1,
      user: { id: 1, name: 'احمد رضایی', email: 'ahmad@example.com' },
      product: { id: 1, name: 'گوشی شیائومی Redmi Note 12', code: 'XMI-001' },
      type: 'stock',
      status: 'pending',
      created_at: '2024-01-15T10:30:00Z'
    },
    {
      id: 2,
      user: { id: 2, name: 'فاطمه احمدی', email: 'fateme@example.com' },
      product: { id: 2, name: 'لپ‌تاپ ایسوس VivoBook', code: 'ASUS-002' },
      type: 'discount',
      status: 'completed',
      created_at: '2024-01-14T14:20:00Z'
    }
  ])

  const topProducts = ref([
    { id: 1, name: 'گوشی شیائومی Redmi Note 12', requests_count: 45 },
    { id: 2, name: 'لپ‌تاپ ایسوس VivoBook', requests_count: 38 },
    { id: 3, name: 'هدفون بلوتوثی', requests_count: 29 },
    { id: 4, name: 'ساعت هوشمند', requests_count: 22 },
    { id: 5, name: 'پاور بانک شیائومی', requests_count: 18 }
  ])

  // متغیرهای تب موجودی
  const stockSearchQuery = ref('')
  const selectedStockItems = ref<number[]>([])
  const stockNotificationStates = ref<Record<number, { isLoading: boolean, error: string | null, success: boolean }>>({})

  const stockRequests = ref([
    {
      id: 1,
      date: '1404/03/14',
      time: '01:43',
      price: 749000,
      phone: '09373722175',
      product: 'ذره بین چراغ قوه در شیائومی مدل XIAODA XD-FDJ01',
      status: 'در انتظار'
    },
    {
      id: 2,
      date: '1404/03/14',
      time: '00:41',
      price: 1960000,
      phone: '09151604766',
      product: 'محلول کن شیائومی شیائومی Deerma NU05 - اصالت و سلامت فیزیکی کالا',
      status: 'در انتظار'
    }
  ])

  // متغیرهای تب تخفیف
  const discountSearchQuery = ref('')
  const selectedDiscountItems = ref<number[]>([])
  const discountNotificationStates = ref<Record<number, { isLoading: boolean, error: string | null, success: boolean }>>({})

  const discountRequests = ref([
    {
      id: 1,
      date: '1404/03/14',
      time: '01:43',
      price: 749000,
      phone: '09373722175',
      product: 'ذره بین چراغ قوه در شیائومی مدل XIAODA XD-FDJ01',
      status: 'در انتظار'
    }
  ])

  // لیست آیتم‌های اطلاع‌رسانی شده
  const stockCompletedRequests = ref<StockRequest[]>([])
  const discountCompletedRequests = ref<StockRequest[]>([])
  const selectedStockCompletedItems = ref<number[]>([])
  const selectedDiscountCompletedItems = ref<number[]>([])

  // Computed properties
  const filteredRequests = computed(() => {
    let filtered = requests.value

    if (searchQuery.value) {
      const query = searchQuery.value.toLowerCase()
      filtered = filtered.filter(req =>
        req.user.name.toLowerCase().includes(query) ||
        req.product.name.toLowerCase().includes(query) ||
        req.user.email.toLowerCase().includes(query)
      )
    }

    if (statusFilter.value) {
      filtered = filtered.filter(req => req.status === statusFilter.value)
    }

    if (typeFilter.value) {
      filtered = filtered.filter(req => req.type === typeFilter.value)
    }

    return filtered
  })

  // Methods
  const exportToExcel = () => {
    // TODO: پیاده‌سازی خروجی اکسل سمت کلاینت یا دریافت از API
    console.warn('Exporting to Excel...')
  }

  const refreshData = async () => {
    // دریافت داده‌ها از API و کش SWR
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    const [{ data: stock }, { data: discount }] = await Promise.all([
      useFetch(`${base}/api/admin/notifications/stock`, { key: 'notify-stock', defaultCache: true, watch: false, server: false, retry: 1, cache: 'force-cache' }),
      useFetch(`${base}/api/admin/notifications/discount`, { key: 'notify-discount', defaultCache: true, watch: false, server: false, retry: 1, cache: 'force-cache' })
    ])
    if (Array.isArray(stock.value)) {
      stockRequests.value = stock.value.map((n: ApiNotification) => ({
        id: n.id,
        date: new Date(n.created_at).toLocaleDateString('fa-IR'),
        time: new Date(n.created_at).toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' }),
        price: 0,
        phone: n.phone || '',
        product: n.product?.name || `#${n.product_id}`,
        status: n.status === 'sent' ? 'اطلاع‌رسانی شده' : n.status === 'pending' ? 'در انتظار' : 'لغو شده'
      }))
    }
    if (Array.isArray(discount.value)) {
      discountRequests.value = discount.value.map((n: ApiNotification) => ({
        id: n.id,
        date: new Date(n.created_at).toLocaleDateString('fa-IR'),
        time: new Date(n.created_at).toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' }),
        price: 0,
        phone: n.phone || '',
        product: n.product?.name || `#${n.product_id}`,
        status: n.status === 'sent' ? 'اطلاع‌رسانی شده' : n.status === 'pending' ? 'در انتظار' : 'لغو شده'
      }))
    }
  }

  const getNavItemCount = (navId: string) => {
    switch (navId) {
      case 'stock':
        return stockRequests.value.length
      case 'discount':
        return discountRequests.value.length
      case 'stock-completed':
        return stockCompletedRequests.value.length
      case 'discount-completed':
        return discountCompletedRequests.value.length
      default:
        return null
    }
  }

  const toggleAllSelection = (event: Event) => {
    const target = event.target as HTMLInputElement
    if (target.checked) {
      selectedRequests.value = filteredRequests.value.map(req => req.id)
    } else {
      selectedRequests.value = []
    }
  }

  const sendBulkNotification = () => {
    console.warn('Sending bulk notifications for:', selectedRequests.value)
  }

  const sendNotification = (request: NotificationRequest) => {
    console.warn('Sending notification for request:', request.id)
  }

  const viewDetails = (request: NotificationRequest) => {
    console.warn('Viewing details for request:', request.id)
  }

  const saveSettings = () => {
    console.warn('Saving settings:', settings.value)
  }

  const saveAdvancedSettings = () => {
    showSettings.value = false
    console.warn('Saving advanced settings:', settings.value)
  }

  const getTypeClass = (type: RequestType) => {
    const classes: Record<RequestType, string> = {
      stock: 'bg-blue-100 text-blue-800',
      discount: 'bg-green-100 text-green-800',
      price: 'bg-purple-100 text-purple-800'
    }
    return classes[type] || 'bg-gray-100 text-gray-800'
  }

  const getTypeLabel = (type: RequestType) => {
    const labels: Record<RequestType, string> = {
      stock: 'موجودی',
      discount: 'تخفیف',
      price: 'قیمت'
    }
    return labels[type] || type
  }

  const getStatusClass = (status: RequestStatus | string) => {
    const oldClasses: Record<RequestStatus, string> = {
      pending: 'bg-yellow-100 text-yellow-800',
      completed: 'bg-green-100 text-green-800',
      cancelled: 'bg-red-100 text-red-800'
    }

    switch (status) {
      case 'در انتظار':
        return 'inline-flex px-3 py-1 text-xs font-semibold bg-orange-100 text-orange-800 rounded-full'
      case 'اطلاع‌رسانی شده':
        return 'inline-flex px-3 py-1 text-xs font-semibold bg-green-100 text-green-800 rounded-full'
      case 'خطا':
        return 'inline-flex px-3 py-1 text-xs font-semibold bg-red-100 text-red-800 rounded-full'
      default:
        if (typeof status === 'string' && status in oldClasses) {
          return oldClasses[status as RequestStatus]
        }
        return 'inline-flex px-3 py-1 text-xs font-semibold bg-gray-100 text-gray-800 rounded-full'
    }
  }

  const getStatusLabel = (status: RequestStatus) => {
    const labels: Record<RequestStatus, string> = {
      pending: 'در انتظار',
      completed: 'تکمیل شده',
      cancelled: 'لغو شده'
    }
    return labels[status] || status
  }

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('fa-IR')
  }

  // Stock tab methods
  const applyFilters = () => {
    console.warn('جستجو:', stockSearchQuery.value)
  }

  const toggleAllStock = () => {
    if (selectedStockItems.value.length === stockRequests.value.length) {
      selectedStockItems.value = []
    } else {
      selectedStockItems.value = stockRequests.value.map(item => item.id)
    }
  }

  const bulkDeleteStock = () => {
    console.warn('حذف موارد انتخاب شده:', selectedStockItems.value)
  }

  const bulkNotifyStock = () => {
    console.warn('ارسال اطلاع‌رسانی برای موارد انتخاب شده:', selectedStockItems.value)
  }

  const sendStockNotification = async (itemId: number) => {
    if (!stockNotificationStates.value[itemId]) {
      stockNotificationStates.value[itemId] = { isLoading: false, error: null, success: false }
    }

    const state = stockNotificationStates.value[itemId]
    state.isLoading = true
    state.error = null

    const item = stockRequests.value.find(req => req.id === itemId)

    try {
      const config = useRuntimeConfig()
      const base = config.public.goApiBase
      await $fetch(`${base}/api/admin/notifications/stock/${itemId}/mark-sent`, { method: 'POST' })

      state.success = true
      state.isLoading = false

      if (item) {
        item.status = 'اطلاع‌رسانی شده'
        const now = new Date()
        const notificationDate = now.toLocaleDateString('fa-IR')
        const notificationTime = now.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })

        stockCompletedRequests.value.push({
          ...item,
          notificationDate,
          notificationTime
        })
        const index = stockRequests.value.findIndex(req => req.id === itemId)
        if (index !== -1) {
          stockRequests.value.splice(index, 1)
        }
        delete stockNotificationStates.value[itemId]
      }
    } catch (error) {
      state.error = error instanceof Error ? error.message : 'خطای نامشخص'
      state.isLoading = false

      if (item) {
        item.status = 'خطا'
      }
    }
  }

  // حذف اعلان موجودی
  const deleteStockNotification = async (itemId: number) => {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    await $fetch(`${base}/api/admin/notifications/stock/${itemId}`, { method: 'DELETE' })
    const idx = stockRequests.value.findIndex(req => req.id === itemId)
    if (idx !== -1) stockRequests.value.splice(idx, 1)
  }

  const clearStockError = (itemId: number) => {
    if (stockNotificationStates.value[itemId]) {
      stockNotificationStates.value[itemId].error = null
    }

    const item = stockRequests.value.find(req => req.id === itemId)
    if (item && item.status === 'خطا') {
      item.status = 'در انتظار'
    }
  }

  // Discount tab methods
  const applyDiscountFilters = () => {
    console.warn('جستجو تخفیف:', discountSearchQuery.value)
  }

  const toggleAllDiscount = () => {
    if (selectedDiscountItems.value.length === discountRequests.value.length) {
      selectedDiscountItems.value = []
    } else {
      selectedDiscountItems.value = discountRequests.value.map(item => item.id)
    }
  }

  const bulkDeleteDiscount = () => {
    console.warn('حذف موارد انتخاب شده تخفیف:', selectedDiscountItems.value)
  }

  const bulkNotifyDiscount = () => {
    console.warn('ارسال اطلاع‌رسانی برای موارد انتخاب شده تخفیف:', selectedDiscountItems.value)
  }

  const sendDiscountNotification = async (itemId: number) => {
    if (!discountNotificationStates.value[itemId]) {
      discountNotificationStates.value[itemId] = { isLoading: false, error: null, success: false }
    }

    const state = discountNotificationStates.value[itemId]
    state.isLoading = true
    state.error = null

    const item = discountRequests.value.find(req => req.id === itemId)

    try {
      const config = useRuntimeConfig()
      const base = config.public.goApiBase
      await $fetch(`${base}/api/admin/notifications/discount/${itemId}/mark-sent`, { method: 'POST' })

      state.success = true
      state.isLoading = false

      if (item) {
        item.status = 'اطلاع‌رسانی شده'
        const now = new Date()
        const notificationDate = now.toLocaleDateString('fa-IR')
        const notificationTime = now.toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' })

        discountCompletedRequests.value.push({
          ...item,
          notificationDate,
          notificationTime
        })
        const index = discountRequests.value.findIndex(req => req.id === itemId)
        if (index !== -1) {
          discountRequests.value.splice(index, 1)
        }
        delete discountNotificationStates.value[itemId]
      }
    } catch (error) {
      state.error = error instanceof Error ? error.message : 'خطای نامشخص'
      state.isLoading = false

      if (item) {
        item.status = 'خطا'
      }
    }
  }

  // حذف اعلان تخفیف
  const deleteDiscountNotification = async (itemId: number) => {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    await $fetch(`${base}/api/admin/notifications/discount/${itemId}`, { method: 'DELETE' })
    const idx = discountRequests.value.findIndex(req => req.id === itemId)
    if (idx !== -1) discountRequests.value.splice(idx, 1)
  }

  const clearDiscountError = (itemId: number) => {
    if (discountNotificationStates.value[itemId]) {
      discountNotificationStates.value[itemId].error = null
    }

    const item = discountRequests.value.find(req => req.id === itemId)
    if (item && item.status === 'خطا') {
      item.status = 'در انتظار'
    }
  }

  // Completed tabs methods
  const toggleAllStockCompleted = () => {
    if (selectedStockCompletedItems.value.length === stockCompletedRequests.value.length) {
      selectedStockCompletedItems.value = []
    } else {
      selectedStockCompletedItems.value = stockCompletedRequests.value.map(item => item.id)
    }
  }

  const toggleAllDiscountCompleted = () => {
    if (selectedDiscountCompletedItems.value.length === discountCompletedRequests.value.length) {
      selectedDiscountCompletedItems.value = []
    } else {
      selectedDiscountCompletedItems.value = discountCompletedRequests.value.map(item => item.id)
    }
  }

  const bulkDeleteStockCompleted = () => {
    selectedStockCompletedItems.value.forEach(itemId => {
      const index = stockCompletedRequests.value.findIndex(item => item.id === itemId)
      if (index !== -1) {
        stockCompletedRequests.value.splice(index, 1)
      }
    })
    selectedStockCompletedItems.value = []
  }

  const bulkDeleteDiscountCompleted = () => {
    selectedDiscountCompletedItems.value.forEach(itemId => {
      const index = discountCompletedRequests.value.findIndex(item => item.id === itemId)
      if (index !== -1) {
        discountCompletedRequests.value.splice(index, 1)
      }
    })
    selectedDiscountCompletedItems.value = []
  }

  const bulkResendStockNotification = () => {
    console.warn('ارسال مجدد اطلاع‌رسانی برای موارد انتخاب شده موجودی:', selectedStockCompletedItems.value)
    selectedStockCompletedItems.value = []
  }

  const bulkResendDiscountNotification = () => {
    console.warn('ارسال مجدد اطلاع‌رسانی برای موارد انتخاب شده تخفیف:', selectedDiscountCompletedItems.value)
    selectedDiscountCompletedItems.value = []
  }

  return {
    // Data
    totalRequests,
    pendingRequests,
    completedRequests,
    todayRequests,
    successRate,
    activeTab,
    activeNavItem,
    searchQuery,
    statusFilter,
    typeFilter,
    selectedRequests,
    showSettings,
    activeSettingsTab,
    settings,
    tabs,
    settingsTabs,
    navigationItems,
    requests,
    topProducts,
    stockSearchQuery,
    selectedStockItems,
    stockNotificationStates,
    stockRequests,
    discountSearchQuery,
    selectedDiscountItems,
    discountNotificationStates,
    discountRequests,
    stockCompletedRequests,
    discountCompletedRequests,
    selectedStockCompletedItems,
    selectedDiscountCompletedItems,

    // Computed
    filteredRequests,

    // Methods
    exportToExcel,
    refreshData,
    getNavItemCount,
    toggleAllSelection,
    sendBulkNotification,
    sendNotification,
    viewDetails,
    saveSettings,
    saveAdvancedSettings,
    getTypeClass,
    getTypeLabel,
    getStatusClass,
    getStatusLabel,
    formatDate,
    applyFilters,
    toggleAllStock,
    bulkDeleteStock,
    bulkNotifyStock,
    sendStockNotification,
    deleteStockNotification,
    clearStockError,
    applyDiscountFilters,
    toggleAllDiscount,
    bulkDeleteDiscount,
    bulkNotifyDiscount,
    sendDiscountNotification,
    deleteDiscountNotification,
    clearDiscountError,
    toggleAllStockCompleted,
    toggleAllDiscountCompleted,
    bulkDeleteStockCompleted,
    bulkDeleteDiscountCompleted,
    bulkResendStockNotification,
    bulkResendDiscountNotification
  }
}