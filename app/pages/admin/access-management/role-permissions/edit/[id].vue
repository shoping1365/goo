<template>
  
  <div class="min-h-screen">
    <!-- Header -->
    <div class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-30">
      <div class="px-6 py-4">
        <div class="flex items-center justify-between">
          <div>
            <h1 class="text-2xl font-bold text-gray-900">ویرایش مجوزهای نقش<span v-if="roleName"> : {{ roleName }}</span></h1>
            <!-- حذف توضیح نقش -->
          </div>
          
          <div class="flex items-center gap-3">
            <button 
              @click="savePermissions"
              class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-lg text-white bg-blue-600 hover:bg-blue-700 transition-colors"
            >
              <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
              </svg>
              ذخیره تغییرات
            </button>
            
            <NuxtLink 
              to="/admin/access-management/permissions"
              class="inline-flex items-center px-4 py-2 border border-gray-200 rounded-lg bg-white hover:bg-gray-50 transition-all shadow-md"
            >
              <svg class="w-5 h-5 ml-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18"></path>
              </svg>
              <span class="text-gray-700">بازگشت</span>
            </NuxtLink>
          </div>
        </div>
      </div>
    </div>
    
    <div class="max-w-6xl mx-auto py-8">
      <div class="bg-white rounded-2xl shadow-lg border border-gray-200 p-6">
        <h2 class="text-lg font-semibold text-gray-900 mb-4">دسترسی‌ها</h2>
        <div v-for="(group, idx) in sidebarMenu" :key="group.id" class="mb-4 rounded-xl overflow-hidden border border-gray-100 bg-gray-50">
          <button @click="toggleAccordion(idx)" class="w-full flex items-center justify-between px-6 py-4 bg-gradient-to-l from-blue-50 to-white hover:from-blue-100 transition-all focus:outline-none">
            <span class="font-bold text-blue-800 text-base flex items-center gap-2">
              {{ group.title }}
              <span v-if="getGroupPermissionStatus(group)" class="text-xs font-normal text-blue-500 mr-2">{{ getGroupPermissionStatus(group) }}</span>
            </span>
            <svg :class="{'rotate-90': openAccordions[idx]}" class="w-5 h-5 text-blue-500 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>
          <transition name="fade">
            <div v-show="openAccordions[idx]" class="bg-white border-t border-gray-100">
              <table class="w-full text-sm text-gray-700 mt-0">
                <thead>
                  <tr class="bg-blue-50 text-blue-900">
                    <th class="py-2 px-4 text-right font-semibold">زیرمنو</th>
                    <th v-for="perm in defaultPermissions" :key="perm.key" class="py-2 px-2 font-semibold text-center">{{ perm.label }}</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in group.items" :key="item.id" class="hover:bg-blue-50 transition">
                    <td class="py-2 px-4 text-right font-medium">{{ item.title }}</td>
                    <td v-for="perm in defaultPermissions" :key="perm.key" class="py-2 px-2 text-center">
                      <input type="checkbox"
                        :checked="selectedPermissions.has(getPermissionName(item.id, perm.key))"
                        @change="togglePermission(item.id, perm.key)"
                        :disabled="perm.key !== 'read' && !selectedPermissions.has(getPermissionName(item.id, 'read'))"
                        class="accent-blue-600 w-4 h-4 rounded border-gray-300 focus:ring-2 focus:ring-blue-400 transition" />
                    </td>
                  </tr>
                </tbody>
              </table>
              <!-- دکمه‌های انتخاب همه و حذف همه زیر جدول -->
              <div class="flex flex-row-reverse gap-3 items-center px-6 py-2 bg-blue-50 border-t border-blue-100 mt-0 rounded-b-xl justify-start">
                <button @click="selectAllGroup(group)" class="px-3 py-1 rounded bg-green-100 text-green-700 font-semibold hover:bg-green-200 transition">انتخاب همه</button>
                <button @click="deselectAllGroup(group)" class="px-3 py-1 rounded bg-red-50 text-red-600 font-semibold hover:bg-red-100 transition">حذف همه</button>
              </div>
            </div>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

// تعریف definePageMeta برای Nuxt 3
declare const definePageMeta: (meta: { layout?: string }) => void

// تعریف interface برای Permission
interface Permission {
  name: string
  [key: string]: unknown
}

// تعریف interface برای Role Data
interface RoleData {
  display_name?: string
  name?: string
  permissions?: Permission[]
  [key: string]: unknown
}

// تعریف interface برای API Response
interface RoleApiResponse {
  success?: boolean
  data?: RoleData
  display_name?: string
  name?: string
  permissions?: Permission[]
  [key: string]: unknown
}

definePageMeta({ layout: 'admin-main' })

const route = useRoute()
const roleId = route.params.id
const roleName = ref('')

// ساختار کامل منوی سایدبار (مطابق admin-main.vue)
const sidebarMenu = [
  {
    id: 'media-management',
    title: 'مدیریت رسانه',
    items: [
      { id: 'media-library', title: 'کتابخانه رسانه' },
      { id: 'media-upload', title: 'افزودن رسانه جدید' },
      { id: 'image-compression', title: 'تنظیمات تصاویر' },
      { id: 'video-compression', title: 'تنظیمات ویدیو' },
    ]
  },
  {
    id: 'posts',
    title: 'مدیریت نوشته‌ها',
    items: [
      { id: 'add-post', title: 'افزودن نوشته' },
      { id: 'all-posts', title: 'همه نوشته‌ها' },
      { id: 'post-categories', title: 'دسته‌بندی‌ها' },
    ]
  },
  {
    id: 'page-management',
    title: 'مدیریت صفحات',
    items: [
      { id: 'create-page', title: 'ایجاد صفحه جدید' },
      { id: 'all-pages', title: 'همه صفحات' },
      { id: 'page-categories', title: 'دسته بندی صفحات' },
    ]
  },
  {
    id: 'product-management',
    title: 'مدیریت محصولات',
    items: [
      { id: 'brands', title: 'برندها' },
      { id: 'notify-requests', title: 'خبرم کن' },
      { id: 'attributes', title: 'ویژگی‌ها' },
      { id: 'product-categories', title: 'دسته‌بندی‌ها' },
      { id: 'attribute-groups', title: 'گروه ویژگی‌ها' },
      { id: 'qa', title: 'پرسش و پاسخ' },
      { id: 'products', title: 'همه محصولات' },
      { id: 'wishlists', title: 'لیست علاقمندی' },
      { id: 'reviews', title: 'نظرات محصولات' },
      { id: 'products-new', title: 'افزودن محصول جدید' },
      { id: 'products-bulk-edit', title: 'ویرایش کلی محصولات' },
    ]
  },
  {
    id: 'users',
    title: 'مدیریت کاربران',
    items: [
      { id: 'users', title: 'کاربران' },
      { id: 'blacklist', title: 'لیست سیاه' },
      { id: 'user-rating-system', title: 'سیستم رتبه‌بندی کاربران' },
    ]
  },
  {
    id: 'access-management',
    title: 'مدیریت دسترسی‌ها',
    items: [
      { id: 'roles', title: 'نقش‌های کاربری' },
      { id: 'permissions', title: 'مجوزها' },
    ]
  },
  {
    id: 'transactions',
    title: 'مدیریت سفارشات',
    items: [
      { id: 'orders', title: 'کل سفارشات' },
      { id: 'orders-processing', title: 'سفارشات در حال انجام' },
      { id: 'orders-cancelled', title: 'سفارشات کنسل شده' },
      { id: 'orders-shipped', title: 'سفارشات ارسال شده' },
      { id: 'order-surveys', title: 'نظر سنجی بعد از خرید' },
      { id: 'order-surveys2', title: 'نظر سنجی بعد از خرید 2' },
    ]
  },
  {
    id: 'finance',
    title: 'مالی و پرداخت',
    items: [
      { id: 'payments', title: 'پرداخت‌ها' },
      { id: 'payment-gateways', title: 'درگاه‌های پرداخت' },
      { id: 'giftcard', title: 'گیفت‌کارت' },
      { id: 'wallet', title: 'کیف پول' },
      { id: 'tax', title: 'مالیات و تطابق' },
      { id: 'accounting', title: 'اتصال حسابداری/ERP' },
    ]
  },
  {
    id: 'shipping',
    title: 'حمل‌ونقل و لجستیک',
    items: [
      { id: 'shipping-methods', title: 'روش‌های ارسال' },
      { id: 'fleet', title: 'مدیریت ناوگان' },
      { id: 'driver-tracking', title: 'رهگیری رانندگان' },
      { id: 'delivery-slots', title: 'زمان‌بندی تحویل' },
    ]
  },
  {
    id: 'marketing',
    title: 'بازاریابی و فروش',
    items: [
      { id: 'marketing', title: 'بازاریابی و فروش' },
      { id: 'discounts', title: 'کوپن و کمپین' },
      { id: 'integrations', title: 'ترب / ایمالز' },
      { id: 'referral', title: 'برنامه ارجاع' },
      { id: 'loyalty', title: 'برنامه وفاداری' },
      { id: 'ab-testing', title: 'تست A/B' },
      { id: 'recommender', title: 'پیشنهادات هوشمند' },
    ]
  },
  {
    id: 'seo',
    title: 'ابزارهای سئو',
    items: [
      { id: 'SchemaManagement', title: 'مدیریت اسکیما' },
      { id: 'keyword-tracking', title: 'تحلیل و ردیابی کلمات کلیدی' },
      { id: 'content-generation', title: 'تولید محتوا' },
      { id: 'page-analysis', title: 'آنالیز صفحات' },
      { id: 'content-optimization', title: 'بهینه‌سازی محتوا و متا' },
      { id: 'site-structure', title: 'ساختار و ناوبری سایت' },
      { id: 'link-management', title: 'مدیریت لینک‌ها' },
      { id: 'image-optimization', title: 'بهینه‌سازی تصاویر' },
      { id: 'local-seo', title: 'سئو محلی' },
      { id: 'performance', title: 'عملکرد و سرعت' },
      { id: 'user-experience', title: 'تجربه کاربری و ریسپانسیو' },
      { id: 'seo-monitoring', title: 'نظارت و گزارش‌گیری' },
      { id: 'redirects', title: 'تغییر مسیر' },
    ]
  },
  {
    id: 'analytics',
    title: 'تحلیل و گزارش‌ها',
    items: [
      { id: 'reports', title: 'گزارش‌ها و تحلیل‌ها' },
      { id: 'bi', title: 'هوش تجاری و داده‌کاوی' },
    ]
  },
  {
    id: 'support-management',
    title: 'مدیریت پشتیبانی',
    items: [
      { id: 'live-chat', title: 'چت آنلاین' },
      { id: 'tickets', title: 'تیکت‌ها' },
      { id: 'chat-window', title: 'پنجره چت' },
    ]
  },
  {
    id: 'content',
    title: 'مدیریت محتوا',
    items: [
      { id: 'banners', title: 'بنرها و اسلایدر' },
      { id: 'menu-management', title: 'مدیریت منوها' },
      { id: 'header-management', title: 'مدیریت هدر' },
      { id: 'footer-management', title: 'مدیریت فوتر' },
    ]
  },
  {
    id: 'notifications',
    title: 'پیامک و اعلان‌ها',
    items: [
      { id: 'system-settings', title: 'تنظیمات سامانه' },
      { id: 'notification-monitoring', title: 'مانیتورینگ' },
      { id: 'alerts', title: 'اعلان‌ها' },
      { id: 'patterns', title: 'پترن‌ها' },
    ]
  },
  {
    id: 'statistics',
    title: 'آمار',
    items: [
      { id: 'visits', title: 'آمار بازدید' },
      { id: 'sales', title: 'آمار فروش' },
      { id: 'statistics-orders', title: 'آمار سفارش‌ها' },
      { id: 'customers', title: 'آمار مشتریان' },
      { id: 'inventory', title: 'آمار موجودی' },
      { id: 'abandoned-carts', title: 'سبدهای خرید رها شده' },
      { id: 'user-behavior', title: 'آمار رفتار کاربران' },
    ]
  },
  {
    id: 'system-security',
    title: 'امنیت سامانه',
    items: [
      { id: 'traffic', title: 'ترافیک آنلاین' },
      { id: 'blocked', title: 'مسدودی' },
      { id: 'login', title: 'ورود' },
      { id: 'scan', title: 'اسکن' },
      { id: 'firewall', title: 'فایروال' },
    ]
  },
  {
    id: 'settings',
    title: 'تنظیمات',
    items: [
      { id: 'settings', title: 'تنظیمات' },
    ]
  },
]

const defaultPermissions = [
  { key: 'read', label: 'مشاهده' },
  { key: 'create', label: 'ایجاد' },
  { key: 'update', label: 'ویرایش' },
  { key: 'delete', label: 'حذف' },
]

const selectedPermissions = ref<Set<string>>(new Set())
const openAccordions = ref<boolean[]>(sidebarMenu.map(() => false))

// مپ تبدیل آیتم به نام صحیح مجوز دیتابیس
const permissionNameMap: Record<string, string> = {
  // مدیریت رسانه
  'media-library': 'media_library',
  'media-upload': 'media_upload',
  'image-compression': 'media_image_compression',
  'video-compression': 'media_video_compression',

  // مدیریت نوشته‌ها
  'add-post': 'post_add',
  'all-posts': 'post_all',
  'post-categories': 'post_categories',

  // مدیریت صفحات
  'create-page': 'page_create',
  'all-pages': 'page_all',
  'page-categories': 'page_categories',

  // مدیریت محصولات
  'brands': 'product_brands',
  'notify-requests': 'product_notify_requests',
  'attributes': 'product_attributes',
  'product-categories': 'product_categories',
  'attribute-groups': 'product_attribute_groups',
  'qa': 'product_qa',
  'products': 'product_products',
  'wishlists': 'product_wishlists',
  'reviews': 'product_reviews',
  'products-new': 'product_products',
  'products-bulk-edit': 'product_products',

  // مدیریت کاربران
  'users': 'user_users',
  'blacklist': 'user_blacklist',
  'user-rating-system': 'user_rating_system',

  // مدیریت دسترسی‌ها
  'roles': 'role_roles',
  'permissions': 'permission_permissions',

  // مدیریت سفارشات
  'orders': 'order_orders',
  'orders-processing': 'order_processing',
  'orders-cancelled': 'order_cancelled',
  'orders-shipped': 'order_shipped',
  'order-surveys': 'order_surveys',
      'order-surveys2': 'order_surveys2',
    'order-fraud-detection': 'order_fraud_detection',

  // مالی و پرداخت
  'payments': 'finance_payments',
  'payment-gateways': 'finance_payment_gateways',
  'giftcard': 'finance_giftcard',
  'wallet': 'finance_wallet',
  'tax': 'finance_tax',
  'accounting': 'finance_accounting',

  // حمل‌ونقل و لجستیک
  'shipping-methods': 'shipping_methods',
  'fleet': 'shipping_fleet',
  'driver-tracking': 'shipping_driver_tracking',
  'delivery-slots': 'shipping_delivery_slots',

  // بازاریابی و فروش
  'marketing': 'marketing',
  'discounts': 'marketing_discounts',
  'integrations': 'marketing_integrations',
  'referral': 'marketing_referral',
  'loyalty': 'marketing_loyalty',
  'ab-testing': 'marketing_ab_testing',
      'recommender': 'marketing_recommender',
    'fraud_detection': 'marketing_fraud_detection',
      'banners': 'content_banners',

  // ابزارهای سئو
  'SchemaManagement': 'seo_schema_management',
  'keyword-tracking': 'seo_keyword_tracking',
  'content-generation': 'seo_content_generation',
  'page-analysis': 'seo_page_analysis',
  'content-optimization': 'seo_content_optimization',
  'site-structure': 'seo_site_structure',
  'link-management': 'seo_link_management',
  'image-optimization': 'seo_image_optimization',
  'local-seo': 'seo_local_seo',
  'performance': 'seo_performance',
  'user-experience': 'seo_user_experience',
  'seo-monitoring': 'seo_monitoring',
  'redirects': 'seo_redirects',

  // تحلیل و گزارش‌ها
  'reports': 'analytics_reports',
  'bi': 'analytics_bi',

  // مدیریت پشتیبانی
  'live-chat': 'support_live_chat',
  'tickets': 'support_tickets',
  'chat-window': 'support_chat_window',

  // مدیریت محتوا
  'menu-management': 'content_menu_management',
  'header-management': 'header_manage',
  'footer-management': 'footer_manage',

  // پیامک و اعلان‌ها
  'system-settings': 'notification_system_settings',
  'notification-monitoring': 'notification_monitoring',
  'alerts': 'notification_alerts',
  'patterns': 'notification_patterns',

  // آمار
  'visits': 'statistics_visits',
  'sales': 'statistics_sales',
  'statistics-orders': 'statistics_orders',
  'customers': 'statistics_customers',
  'inventory': 'statistics_inventory',
  'abandoned-carts': 'statistics_abandoned_carts',
  'user-behavior': 'statistics_user_behavior',

  // امنیت سامانه
  'traffic': 'security_traffic',
  'blocked': 'security_blocked',
  'login': 'security_login',
  'scan': 'security_scan',
  'firewall': 'security_firewall',

  // تنظیمات
  'settings': 'settings',
}

// تابع ساخت نام مجوز صحیح
function getPermissionName(itemId: string, permKey: string) {
  const base = permissionNameMap[itemId] || itemId
  // تبدیل dash به underline
  return (base + '.' + permKey).replace(/-/g, '_')
}

onMounted(async () => {
  try {
    const response = await $fetch<RoleApiResponse>(`/api/roles/${roleId}`)
    // بررسی ساختار response
    const roleData: RoleData = response.data || response
    roleName.value = roleData.display_name || roleData.name || 'نقش نامشخص'
    if (roleData.permissions && Array.isArray(roleData.permissions)) {
      selectedPermissions.value = new Set(roleData.permissions.map((p: Permission) => p.name))
    }
  } catch (e) {
    console.error('خطا در دریافت اطلاعات نقش:', e)
    roleName.value = 'نقش نامشخص'
  }
})

function toggleAccordion(idx: number) {
  openAccordions.value[idx] = !openAccordions.value[idx]
}

function togglePermission(itemId: string, permKey: string) {
  const permName = getPermissionName(itemId, permKey)
  if (selectedPermissions.value.has(permName)) {
    selectedPermissions.value.delete(permName)
  } else {
    selectedPermissions.value.add(permName)
  }
}

// محاسبه تعداد مجوزهای فعال هر گروه
function getGroupPermissionStatus(group: any) {
  const total = group.items.length * defaultPermissions.length
  let count = 0
  for (const item of group.items) {
    for (const perm of defaultPermissions) {
      if (selectedPermissions.value.has(getPermissionName(item.id, perm.key))) count++
    }
  }
  if (count === 0) return ''
  if (count === total) return 'کامل'
  return count + ' مجوز'
}

// توابع انتخاب همه و حذف همه برای هر گروه
function selectAllGroup(group: any) {
  for (const item of group.items) {
    for (const perm of defaultPermissions) {
      selectedPermissions.value.add(getPermissionName(item.id, perm.key))
    }
  }
}
function deselectAllGroup(group: any) {
  for (const item of group.items) {
    for (const perm of defaultPermissions) {
      selectedPermissions.value.delete(getPermissionName(item.id, perm.key))
    }
  }
}

async function savePermissions() {
  try {
    // هنگام ارسال، همه نام‌ها را با آندرلاین بفرست
    const permissionsToSend = Array.from(selectedPermissions.value).map(p => p.replace(/-/g, '_'))
    await $fetch(`/api/roles/${roleId}/permissions`, {
      method: 'PUT',
      body: { permissions: permissionsToSend },
      headers: { 'Content-Type': 'application/json' }
    })
    alert('مجوزها با موفقیت ذخیره شدند')
  } catch (e) {
    alert('خطا در ذخیره مجوزها')
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style> 