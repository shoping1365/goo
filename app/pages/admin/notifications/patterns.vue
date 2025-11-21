    <template>
      <div>
        <!-- Page Header -->
        <div class="shadow-sm border-b border-gray-200 mb-4 bg-white">
          <div class="w-full px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between items-center py-3">
              <div>
                <h1 class="text-lg font-bold text-gray-900">مدیریت پترن‌ها</h1>
                <p class="text-xs text-gray-600 mt-1">مدیریت الگوها و قالب‌های پیامک و اعلان‌ها</p>
              </div>
              <div class="flex space-x-2 space-x-reverse">
                <button
                    class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-gradient-to-r from-purple-400 to-purple-600 hover:from-purple-500 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 shadow-md transition-all duration-200 hover:shadow-lg hover:scale-105"
                    @click="showCreateModal = true"
                >
                  <svg class="w-5 h-5 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
                  </svg>
                  ایجاد پترن جدید
                </button>
              </div>
            </div>
          </div>
        </div>

        <div class="w-full px-4 py-2">
          <div class="mx-auto px-4 sm:px-6 lg:px-8 py-4">
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3 mb-4">
                        <TemplateCard 
                title="کل پترن‌ها" 
                :value="stats.totalPatterns" 
                variant="blue"
                @click="setStatusFilter('')"
              >
                <template #icon>
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                </template>
              </TemplateCard>

                        <TemplateCard 
                title="پترن‌های فعال" 
                :value="stats.activePatterns" 
                variant="green"
                @click="setStatusFilter('active')"
              >
                <template #icon>
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
                  </svg>
                </template>
              </TemplateCard>

                        <TemplateCard 
                title="استفاده امروز" 
                :value="stats.todayUsage" 
                variant="orange"
                @click="setStatusFilter('inactive')"
              >
                <template #icon>
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                  </svg>
                </template>
              </TemplateCard>

                                  <TemplateCard 
              title="نرخ موفقیت" 
              :value="stats.successRate + '%'" 
              variant="cyan"
              @click="setStatusFilter('draft')"
            >
                <template #icon>
                  <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                  </svg>
                </template>
              </TemplateCard>
            </div>

          <!-- فیلترها -->
          <div class="shadow-lg rounded-2xl mb-4 border border-blue-100 bg-white">
            <div class="px-4 py-3 border-b border-blue-100 flex items-center gap-2">
              <svg class="w-5 h-5 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4a1 1 0 011-1h16a1 1 0 011 1v2.586a1 1 0 01-.293.707l-6.414 6.414a1 1 0 00-.293.707V17l-4 4v-6.586a1 1 0 00-.293-.707L3.293 7.707A1 1 0 013 7V4z"></path></svg>
              <h3 class="text-base font-bold text-blue-900">فیلترها و جستجو</h3>
            </div>
            <div class="p-6">
              <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">جستجو در پترن‌ها</label>
                  <div class="relative">
                                                    <input
                      v-model="filters.search"
                      type="text"
                      class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-blue-500 focus:border-blue-500 text-sm"
                      placeholder="نام پترن یا محتوا..."
                      dir="rtl"
                  />
                    <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                      <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                      </svg>
                    </div>
                  </div>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">نوع پترن</label>
                                <select v-model="filters.type" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                    <option value="">همه انواع</option>
                    <option value="sms">پیامک</option>
                    <option value="email">ایمیل</option>
                    <option value="push">اعلان مرورگر</option>
                  </select>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">دسته‌بندی</label>
                                <select v-model="filters.category" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                    <option value="">همه دسته‌ها</option>
                    <option value="order">سفارش</option>
                    <option value="marketing">بازاریابی</option>
                    <option value="notification">اعلان</option>
                    <option value="verification">تایید</option>
                  </select>
                </div>

                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">درگاه پیامکی</label>
                                <select v-model="filters.gateway" class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-sm">
                    <option value="">همه درگاه‌ها</option>
                                    <option v-for="gateway in gateways" :key="gateway.id" :value="gateway.id.toString()">
                      {{ gateway.name }}
                    </option>
                  </select>
                </div>
              </div>
            </div>
          </div>

                      <!-- کارت‌های پترن گروه‌بندی شده -->
        <div v-for="gateway in groupedPatterns" :key="gateway.id" class="mb-4 bg-white rounded-xl shadow-lg border border-gray-200 p-6">
            
            <!-- هدر درگاه -->
            <div class="flex items-center justify-between mb-4">
              <div class="flex items-center space-x-3 space-x-reverse">
                <div class="w-8 h-8 rounded-lg flex items-center justify-center">
                  <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
                  </svg>
                </div>
                <div>
                  <h2 class="text-xl font-bold text-gray-900">{{ gateway.name }}</h2>
                  <p class="text-sm text-gray-600">{{ gateway.patterns.length }} پترن</p>
                </div>
              </div>
              <div class="flex items-center space-x-2 space-x-reverse">
                <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800">
                  {{ gateway.activeCount }} فعال
                </span>
                <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-blue-100 text-blue-800">
                  {{ gateway.totalUsage }} استفاده
                </span>
              </div>
            </div>

            <!-- کارت‌های پترن -->
            <div v-if="gateway.patterns.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
              <PatternCard 
                v-for="pattern in gateway.patterns" 
                :key="`card-${pattern.id}`" 
                :pattern="pattern"
                @edit="editPattern"
                @duplicate="duplicatePattern"
                @test="testPattern"
                @delete="deletePattern"
              />
            </div>

            <!-- پیام خالی برای درگاه -->
            <div v-else class="text-center py-8 rounded-lg border-2 border-dashed border-gray-300 bg-gray-50">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
              </svg>
              <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ پترنی برای این درگاه وجود ندارد</h3>
              <p class="mt-1 text-sm text-gray-500">اولین پترن را برای {{ gateway.name }} ایجاد کنید</p>
            </div>
          </div>

                        <!-- پیام خالی -->
        <div v-if="groupedPatterns.length === 0" class="text-center py-12">
              <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"></path>
              </svg>
              <h3 class="mt-2 text-sm font-medium text-gray-900">هیچ درگاه پیامکی یافت نشد</h3>
              <p class="mt-1 text-sm text-gray-500">ابتدا یک درگاه پیامکی ایجاد کنید</p>
              <div class="mt-6">
                          <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700" @click="showCreateModal = true">
                  ایجاد درگاه جدید
              </button>
            </div>
          </div>
        </div>
        </div>

        <!-- Modal ایجاد/ویرایش پترن -->
        <PatternForm 
          :show="showCreateModal"
          :editing-pattern="editingPattern"
          @close="closeModal"
          @save="handleSavePattern"
        />

                          <!-- Modal تست پترن -->
    <div v-if="showTestModal" class="fixed inset-0 z-50">
      <div class="flex items-center justify-center min-h-screen p-6">
        <div class="relative bg-white rounded-lg shadow-xl max-w-md w-full mx-auto border-2 border-blue-500">
          <div class="px-6 py-4">
            <h3 class="text-lg font-medium text-gray-900 mb-4 text-center">تست پترن</h3>
            
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">شماره موبایل</label>
                <input 
                  v-model="testData.phone" 
                  type="text" 
                  placeholder="شماره موبایل" 
                  class="block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50"
                >
              </div>
              
              <!-- فیلدهای متغیرها -->
              <div v-if="getPatternVariables().length > 0">
                <label class="block text-sm font-medium text-gray-700 mb-2">متغیرها</label>
                <div class="space-y-2">
                  <div v-for="variable in getPatternVariables()" :key="variable" class="flex items-center space-x-2 space-x-reverse">
                    <label class="text-sm text-gray-600 min-w-[80px]">{{ variable }}:</label>
                    <input 
                      v-if="!['gateway', 'date', 'time'].includes(variable)"
                      v-model="testData.variables[variable]" 
                      type="text" 
                      :placeholder="variable"
                      class="flex-1 rounded-md border-gray-300 shadow-sm focus:border-blue-300 focus:ring focus:ring-blue-200 focus:ring-opacity-50 text-sm"
                    >
                    <span v-else class="flex-1 text-sm text-gray-500 bg-gray-100 px-2 py-1 rounded">
                      {{ variable === 'gateway' ? testingPattern?.gatewayName : 
                         variable === 'date' ? new Date().toLocaleDateString('fa-IR') : 
                         new Date().toLocaleTimeString('fa-IR') }}
                    </span>
                  </div>
                </div>
              </div>
              
              <div class="rounded-lg p-3 bg-blue-100 border border-blue-200">
                <p class="text-sm text-gray-700 font-medium mb-2">پیش‌نمایش:</p>
                <p class="text-sm text-gray-600">{{ getPreviewContent() }}</p>
              </div>
            </div>
          </div>
          
          <div class="bg-gray-50 px-6 py-3 flex justify-end space-x-8 space-x-reverse">
            <button 
              class="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 ml-4" 
              @click="sendTestMessage"
            >
              ارسال تست
            </button>
            <button 
              class="px-4 py-2 bg-gray-300 text-gray-700 rounded-md hover:bg-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-500" 
              @click="showTestModal = false"
            >
              انصراف
            </button>
          </div>
        </div>
      </div>
    </div>
      </div>
    </template>

    <script setup lang="ts">
    // @ts-ignore: Nuxt macro
    definePageMeta({
      layout: 'admin-main',
      middleware: 'admin'
    })

      // Import دستی کامپوننت PatternCard
    import { computed, onMounted, ref, watch } from 'vue'
import PatternCard from '~/components/admin/sms/PatternCard.vue'
import PatternForm from '~/components/common/PatternForm.vue'
import TemplateCard from '~/components/common/TemplateCard.vue'

    // تعریف interface ها
    interface Pattern {
      id: number
      name: string
      description: string
      type: 'sms' | 'email' | 'push'
      category: 'order' | 'marketing' | 'notification' | 'verification'
      scope: 'customer' | 'manager'
      feature: string
      content: string
      status: 'active' | 'inactive' | 'draft'
      usageCount: number
      maxLength?: number
      gatewayId: number
      gatewayName: string
      patternCode: string
      variables: string
    }

    interface Gateway {
      id: number
      name: string
      status: 'active' | 'inactive'
      api_key?: string
      username?: string
      password?: string
    }

    interface PatternForm {
      fixedId?: number
      name: string
      description: string
      type: 'sms' | 'email' | 'push'
      category: 'order' | 'marketing' | 'notification' | 'verification'
      scope: 'customer' | 'manager'
      feature: string
      content: string
      status: 'active' | 'inactive' | 'draft'
      maxLength?: number
      gatewayId: number
      gatewayName: string
      patternCode: string
      variables: string
    }

    interface Filters {
      type: string
      category: string
      status: string
      search: string
      gateway: string
    }

    interface Stats {
      totalPatterns: number
      activePatterns: number
      todayUsage: number
      successRate: number
    }

      interface TestData {
      phone: string
      variables: Record<string, string>
    }

    interface GatewayItem {
      id: number
      name: string
      is_active?: boolean
      api_url?: string
      username?: string
      password?: string
      [key: string]: unknown
    }

    interface PatternItem {
      id: number
      name: string
      description?: string
      scope?: string
      feature?: string
      message_template?: string
      content?: string
      status?: string
      usage_count?: number
      gateway_id?: number
      gateway?: { name?: string }
      pattern_code?: string
      variables?: string
      [key: string]: unknown
    }

    interface APIResponse {
      success?: boolean
      status?: string
      data: GatewayItem[] | PatternItem[] | unknown
      message?: string
    }

    // متغیرهای reactive
    const patterns = ref<Pattern[]>([])
    const gateways = ref<Gateway[]>([])

    // بارگذاری درگاه‌های پیامکی از API
    const loadGateways = async () => {
      try {
        const response = await $fetch('/api/sms-gateways/gateways') as APIResponse
        if (response.status === 'success' && Array.isArray(response.data)) {
          gateways.value = (response.data as GatewayItem[]).map((item: GatewayItem) => ({
            id: item.id,
            name: item.name,
            status: item.is_active ? 'active' as const : 'inactive' as const,
            api_key: item.api_url,
            username: item.username,
            password: item.password
          }))
        } else {
          // اگر داده‌ای دریافت نشد، آرایه خالی تنظیم کن
          gateways.value = []
        }
      } catch {
        gateways.value = []
      }
    }

    // بارگذاری پترن‌ها از API
    const loadPatterns = async () => {
      try {
        const response = await $fetch('/api/admin/sms-patterns') as APIResponse
        
        // بررسی موفقیت با شرط‌های مختلف
        if (response.success || response.status === 'success' || (response.data && Array.isArray(response.data) && response.data.length > 0)) {
          patterns.value = (response.data as PatternItem[]).map((item: PatternItem) => {
            return {
              id: item.id,
              name: item.name,
              description: item.description || '',
              type: 'sms' as const,
              category: 'order' as const,
              scope: (item.scope === 'customer' || item.scope === 'manager') ? (item.scope as 'customer' | 'manager') : 'customer' as 'customer' | 'manager',
              feature: item.feature || '',
              content: item.message_template || item.content || '',
              status: item.status === 'active' ? 'active' as const : 'inactive' as const,
              usageCount: item.usage_count || 0,
              maxLength: 160,
              gatewayId: item.gateway_id || 1,
              gatewayName: item.gateway?.name || 'IPPanel',
              patternCode: item.pattern_code || '',
              variables: item.variables || ''
            }
          })
        } else {
          // اگر داده‌ای دریافت نشد، آرایه خالی تنظیم کن
          patterns.value = []
        }
      } catch {
        patterns.value = []
      }
    }

    // بارگذاری آمار از API
    const loadStats = async () => {
      try {
        interface StatsResponse {
          total_patterns: number
          active_patterns: number
          inactive_patterns: number
          draft_patterns: number
          total_usage: number
        }
        const response = await $fetch<{success?: boolean, data: StatsResponse}>('/api/admin/sms-patterns/stats')
        if (response.success) {
          apiStats.value = response.data
        }
      } catch {
        // خطا در بارگذاری آمار
      }
    }

    // به‌روزرسانی stats
    const updateStats = () => {
      try {
        if (apiStats.value) {
          const successRate = apiStats.value.total_patterns > 0
            ? Math.round((apiStats.value.active_patterns / apiStats.value.total_patterns) * 100)
            : 0
          stats.value.totalPatterns = apiStats.value.total_patterns
          stats.value.activePatterns = apiStats.value.active_patterns
          stats.value.todayUsage = apiStats.value.total_usage
          stats.value.successRate = successRate
        } else {
          // استفاده از مقادیر پیش‌فرض
          stats.value.totalPatterns = 0
          stats.value.activePatterns = 0
          stats.value.todayUsage = 0
          stats.value.successRate = 0
        }
      } catch {
        // خطا در به‌روزرسانی آمار
        stats.value.totalPatterns = 0
        stats.value.activePatterns = 0
        stats.value.todayUsage = 0
        stats.value.successRate = 0
      }
    }

      // بارگذاری داده‌ها در ابتدای صفحه
    onMounted(() => {
      loadGateways()
      loadPatterns()
      loadStats()
      loadAdminPhone() // بارگذاری شماره موبایل مدیر
      updateStats() // تنظیم آمار اولیه
    })

      // Watcher برای تغییرات
    watch(patterns, () => {}, { deep: true })
    watch(gateways, () => {}, { deep: true })
  
    // بارگذاری شماره موبایل مدیر از تنظیمات فروشگاه
    const loadAdminPhone = async () => {
      try {
        interface ShopSettingsResponse {
          status?: string
          data?: {
            adminPhones?: string[]
            [key: string]: unknown
          }
        }
        const response = await $fetch<ShopSettingsResponse>('/api/admin/shop-settings')
        if (response.status === 'success' && response.data?.adminPhones && response.data.adminPhones.length > 0) {
          testData.value.phone = response.data.adminPhones[0]
        }
      } catch {
        // خطا در بارگذاری شماره موبایل مدیر
      }
    }

    // متغیر برای ذخیره آمار از API
    const apiStats = ref<{
      total_patterns: number
      active_patterns: number
      inactive_patterns: number
      draft_patterns: number
      total_usage: number
    } | null>(null)

    const stats = ref<Stats>({
      totalPatterns: 0,
      activePatterns: 0,
      todayUsage: 0,
      successRate: 0
    })

    const filters = ref<Filters>({
      type: '',
      category: '',
      status: '',
      search: '',
      gateway: ''
    })

    const showCreateModal = ref(false)
    const showTestModal = ref(false)
    const editingPattern = ref<Pattern | null>(null)
    const testingPattern = ref<Pattern | null>(null)

    const patternForm = ref<PatternForm>({
      fixedId: undefined,
      name: '',
      description: '',
      type: 'sms',
      category: 'order',
      scope: 'customer',
      feature: '',
      content: '',
      status: 'active',
      maxLength: 160,
      gatewayId: 1,
      gatewayName: 'IPPanel',
      patternCode: '',
      variables: ''
    })

      const testData = ref<TestData>({
      phone: '', // شماره موبایل خالی
      variables: {} as Record<string, string>
    })

    // Computed properties
    const filteredPatterns = computed(() => {
      let filtered = patterns.value

      if (filters.value.type) {
        filtered = filtered.filter(pattern => pattern.type === filters.value.type)
      }

      if (filters.value.category) {
        filtered = filtered.filter(pattern => pattern.category === filters.value.category)
      }

      if (filters.value.status) {
        filtered = filtered.filter(pattern => pattern.status === filters.value.status)
      }

      if (filters.value.gateway) {
        filtered = filtered.filter(pattern => pattern.gatewayId.toString() === filters.value.gateway)
      }

      if (filters.value.search) {
        const search = filters.value.search.toLowerCase()
        filtered = filtered.filter(pattern => 
          pattern.name.toLowerCase().includes(search) || 
          pattern.description.toLowerCase().includes(search) ||
          pattern.content.toLowerCase().includes(search)
        )
      }

      return filtered
    })

    const groupedPatterns = computed(() => {
      const groups: { [key: number]: { id: number; name: string; patterns: Pattern[]; activeCount: number; totalUsage: number; status: string } } = {}
      
      // ابتدا همه درگاه‌ها را اضافه می‌کنیم (حتی اگر پترنی نداشته باشند)
      gateways.value.forEach(gateway => {
        groups[gateway.id] = {
          id: gateway.id,
          name: gateway.name,
          patterns: [],
          activeCount: 0,
          totalUsage: 0,
          status: gateway.status
        }
      })
      
      // سپس پترن‌ها را به درگاه‌های مربوطه اضافه می‌کنیم
      filteredPatterns.value.forEach(pattern => {
        if (groups[pattern.gatewayId]) {
          // اگر درگاه وجود دارد، پترن را به آن اضافه کن
          groups[pattern.gatewayId].patterns.push(pattern)
          if (pattern.status === 'active') {
            groups[pattern.gatewayId].activeCount++
          }
          groups[pattern.gatewayId].totalUsage += pattern.usageCount
        } else {
          // اگر درگاه وجود ندارد، یک گروه موقت برای پترن ایجاد کن
          const tempGatewayId = pattern.gatewayId
          groups[tempGatewayId] = {
            id: tempGatewayId,
            name: pattern.gatewayName || `درگاه ${tempGatewayId}`,
            patterns: [pattern],
            activeCount: pattern.status === 'active' ? 1 : 0,
            totalUsage: pattern.usageCount,
            status: 'active'
          }
        }
      })
      
      const result = Object.values(groups).sort((a, b) => a.name.localeCompare(b.name))
      return result
    })

    // Watcher برای groupedPatterns
    watch(groupedPatterns, () => {}, { deep: true })

    // متدها

    const editPattern = (pattern: Pattern) => {
      editingPattern.value = pattern
      showCreateModal.value = true
    }

    const _createPatternForGateway = (gatewayId: number, gatewayName: string) => {
      editingPattern.value = null
      patternForm.value.gatewayId = gatewayId
      patternForm.value.gatewayName = gatewayName
      showCreateModal.value = true
    }

    const closeModal = () => {
      showCreateModal.value = false
      editingPattern.value = null
    }

    const handleSavePattern = async (formData: PatternForm) => {
      try {
        // تبدیل داده‌های فرم به فرمت API
        const apiData = {
          name: formData.name,
          pattern_code: formData.patternCode,
          description: formData.description,
          variables: formData.variables,
          message_template: formData.content,
          gateway_id: formData.gatewayId,
          status: formData.status === 'active' ? 'active' : 'inactive',
          type: formData.type || 'sms',
          scope: formData.scope || 'customer',
          feature: formData.feature || '',
          fixed_id: formData.fixedId
        }

        if (editingPattern.value) {
          // ویرایش پترن موجود
          await $fetch(`/api/admin/sms-patterns/${editingPattern.value.id}`, {
            method: 'PUT',
            body: apiData
          })
          
          // به‌روزرسانی آرایه محلی
          const index = patterns.value.findIndex(p => p.id === editingPattern.value!.id)
          if (index !== -1) {
            patterns.value[index] = {
              ...patterns.value[index],
              name: formData.name,
              description: formData.description,
              type: formData.type,
              category: formData.category,
              scope: formData.scope,
              feature: formData.feature,
              content: formData.content,
              status: formData.status,
              maxLength: formData.maxLength,
              gatewayId: formData.gatewayId,
              gatewayName: formData.gatewayName,
              patternCode: formData.patternCode,
              variables: formData.variables
            }
          }
        } else {
          // ایجاد پترن جدید
          const response = await $fetch('/api/admin/sms-patterns', {
            method: 'POST',
            body: apiData
          })
          
          // اضافه کردن به آرایه محلی
          interface CreatePatternResponse {
            id: number
            [key: string]: unknown
          }
          const createResponse = response as { data: CreatePatternResponse }
          const newPattern: Pattern = {
            id: createResponse.data.id,
            name: formData.name,
            description: formData.description,
            type: formData.type,
            category: formData.category,
            scope: formData.scope,
            feature: formData.feature,
            content: formData.content,
            status: formData.status,
            usageCount: 0,
            maxLength: formData.maxLength,
            gatewayId: formData.gatewayId,
            gatewayName: formData.gatewayName,
            patternCode: formData.patternCode || '',
            variables: formData.variables || ''
          }
          patterns.value.unshift(newPattern)
        }

        showCreateModal.value = false
        editingPattern.value = null
        alert('پترن با موفقیت ذخیره شد')
      } catch (e: unknown) {
        interface ErrorResponse {
          data?: {
            code?: string
            message?: string
          }
        }
        const error = e as ErrorResponse
        if (error?.data?.code === 'DUPLICATE_FEATURE') {
          alert(error.data.message || 'برای این درگاه و هدف قبلاً پترن فعالی وجود دارد')
        } else {
          alert('خطا در ذخیره پترن')
        }
      }
    }

    const duplicatePattern = async (pattern: Pattern) => {
      try {
        // ارسال درخواست به سرور برای کپی کردن پترن
        const response = await $fetch(`/api/admin/sms-patterns/${pattern.id}/duplicate`, {
          method: 'POST'
        })

        // اضافه کردن کپی جدید به آرایه محلی
        const duplicatedResponse = response as { data: PatternItem }
        const duplicatedPattern = duplicatedResponse.data
        const newPattern: Pattern = {
          id: duplicatedPattern.id,
          name: duplicatedPattern.name,
          description: duplicatedPattern.description || '',
          type: pattern.type,
          category: pattern.category,
          scope: pattern.scope,
          feature: pattern.feature,
          content: duplicatedPattern.message_template || duplicatedPattern.content || '',
          status: (duplicatedPattern.status === 'active' ? 'active' : 'inactive') as 'active' | 'inactive',
          usageCount: duplicatedPattern.usage_count || 0,
          maxLength: pattern.maxLength,
          gatewayId: duplicatedPattern.gateway_id || pattern.gatewayId,
          gatewayName: duplicatedPattern.gateway?.name || pattern.gatewayName,
          patternCode: duplicatedPattern.pattern_code || '',
          variables: duplicatedPattern.variables || ''
        }

        patterns.value.unshift(newPattern)
        alert('پترن با موفقیت کپی شد')
      } catch {
        alert('خطا در کپی کردن پترن')
      }
    }

      const testPattern = async (pattern: Pattern) => {
      testingPattern.value = pattern
      
      // بارگذاری شماره موبایل مدیر از تنظیمات
      try {
        interface ShopSettingsResponse {
          status?: string
          data?: {
            adminPhones?: string[]
            [key: string]: unknown
          }
        }
        const response = await $fetch<ShopSettingsResponse>('/api/admin/shop-settings')
        if (response.status === 'success' && response.data?.adminPhones && response.data.adminPhones.length > 0) {
          testData.value.phone = response.data.adminPhones[0]
        } else {
          testData.value.phone = '' // شماره موبایل خالی
        }
      } catch {
        testData.value.phone = '' // شماره موبایل خالی
      }
      
      showTestModal.value = true
    }





    const deletePattern = async (pattern: Pattern) => {
      if (confirm('آیا مطمئن هستید که می‌خواهید این پترن را حذف کنید؟')) {
        try {
          await $fetch(`/api/admin/sms-patterns/${pattern.id}`, {
            method: 'DELETE'
          })
          
          patterns.value = patterns.value.filter(p => p.id !== pattern.id)
          alert('پترن با موفقیت حذف شد')
              } catch {
        alert('خطا در حذف پترن')
      }
      }
    }



             const getPreviewContent = () => {
       if (!testingPattern.value) return ''
       
       let content = testingPattern.value.content
       
       // جایگزینی متغیرها در متن
       if (testingPattern.value.variables && testData.value.variables) {
         Object.keys(testData.value.variables).forEach(variable => {
           const placeholder = `%${variable}%`
           const value = testData.value.variables[variable] || `[${variable}]`
           content = content.replace(new RegExp(placeholder, 'g'), value)
         })
       }
       
       // جایگزینی متغیرهای سیستمی
       const now = new Date()
       const gateway = testingPattern.value.gatewayName || 'IPPanel'
       
       content = content.replace(/%gateway%/g, gateway)
       content = content.replace(/%date%/g, now.toLocaleDateString('fa-IR'))
       content = content.replace(/%time%/g, now.toLocaleTimeString('fa-IR'))
       
       return content
     }

         const getPatternVariables = () => {
       if (!testingPattern.value?.variables) return []
       
       // استخراج متغیرها از رشته variables (مثل "gateway,date" یا "%code%,%name%")
       const vars = testingPattern.value.variables.split(',')
       const patternVars = vars
         .map(v => v.trim())
         .filter(v => v !== '')
         .map(v => {
           // اگر متغیر با % شروع و پایان می‌شود، % را حذف کن
           if (v.startsWith('%') && v.endsWith('%')) {
             return v.slice(1, -1)
           }
           // در غیر این صورت، همان متغیر را برگردان
           return v
         })
       
       // اضافه کردن متغیرهای سیستمی که همیشه باید موجود باشند
       const systemVars = ['gateway', 'date', 'time']
       const allVars = [...new Set([...patternVars, ...systemVars])]
       
       return allVars
     }

    const setStatusFilter = (status: string) => {
      filters.value.status = status
    }

         const sendTestMessage = async () => {
       try {
         if (!testData.value.phone) {
           alert('لطفاً شماره موبایل را وارد کنید')
           return
         }

         if (!testingPattern.value) {
           alert('پترن انتخاب نشده')
           return
         }

         // آماده‌سازی متغیرها برای ارسال
         const variables: Record<string, string | number> = {}
         
         // استخراج متغیرهای پترن از رشته variables
         if (testingPattern.value.variables) {
           const patternVars = testingPattern.value.variables
             .split(',')
             .map(v => v.trim())
             .filter(v => v !== '')
             .map(v => {
               // اگر متغیر با % شروع و پایان می‌شود، % را حذف کن
               if (v.startsWith('%') && v.endsWith('%')) {
                 return v.slice(1, -1)
               }
               // در غیر این صورت، همان متغیر را برگردان
               return v
             })
           
           // برای هر متغیر پترن، مقدار را ارسال کن
           patternVars.forEach(variable => {
             const value = testData.value.variables[variable] || ''
             variables[variable] = value
           })
         }
         
         // اضافه کردن متغیرهای سیستمی
         const now = new Date()
         const gateway = testingPattern.value.gatewayName || 'IPPanel'
         
         variables['gateway'] = gateway
         variables['date'] = now.toLocaleDateString('fa-IR')
         variables['time'] = now.toLocaleTimeString('fa-IR')

         const requestData = {
           phone: testData.value.phone,
           variables: variables
         }

         // API call برای ارسال پیام تست
         await $fetch(`/api/admin/sms-patterns/${testingPattern.value.id}/test`, {
           method: 'POST',
           body: requestData
         })

         showTestModal.value = false
         alert('پیام تست با موفقیت ارسال شد')
       } catch {
         alert('خطا در ارسال پیام تست')
       }
     }
    </script>
