<template>
  <div class="footer-trust-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- عنوان بخش -->
      <h4 v-if="title" class="footer-widget-text font-semibold mb-4 text-lg">{{ title }}</h4>
      
      <!-- توضیحات -->
      <p v-if="description" class="footer-widget-text text-opacity-80 mb-6 text-sm leading-relaxed">
        {{ description }}
      </p>
      
      <!-- نشان‌های اعتماد با کد HTML -->
      <div v-if="trustBadges && trustBadges.length > 0" class="space-y-4">
        <div class="grid grid-cols-2 gap-6">
          <div
            v-for="badge in trustBadges"
            :key="badge.id"
            class="bg-gray-50 rounded-lg p-3 border border-gray-200 hover:bg-opacity-20 transition-all duration-200"
          >
            <!-- رندر کد HTML نماد -->
            <!-- 
              ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
              
              این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
              
              ✅ راه حل صحیح:
              1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
              2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
              3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
              
              مثال صحیح:
              import DOMPurify from 'dompurify'
              const sanitizedHtmlCode = computed(() => DOMPurify.sanitize(badge.htmlCode))
              <div v-html="sanitizedHtmlCode"></div>
            -->
            <!-- eslint-disable-next-line vue/no-v-html -->
            <div v-if="badge.htmlCode" class="trust-badge-html" v-html="badge.htmlCode"></div>
            
            <!-- اگر کد نداشت، نمایش پیش‌فرض -->
            <template v-else>
              <!-- آیکون -->
              <div class="text-3xl mb-2">{{ badge.icon || '🏆' }}</div>
              
              <!-- عنوان -->
              <div class="footer-widget-text font-medium text-sm mb-1">{{ badge.title }}</div>
              
              <!-- توضیحات -->
              <div v-if="badge.description" class="footer-widget-text text-opacity-70 text-xs">
                {{ badge.description }}
              </div>
            </template>
          </div>
        </div>
      </div>
      
      <!-- نشان‌های پیش‌فرض -->
      <div v-else class="space-y-4">
        <div class="grid grid-cols-2 gap-6">
          <!-- امنیت -->
          <div class="bg-gray-50 rounded-lg p-3 border border-gray-200">
            <div class="text-3xl mb-2">🔒</div>
            <div class="footer-widget-text font-medium text-sm mb-1">امنیت بالا</div>
            <div class="footer-widget-text text-opacity-70 text-xs">SSL رمزگذاری شده</div>
            <div class="mt-2">
              <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-500 bg-opacity-20 text-green-300">
                <span class="w-2 h-2 rounded-full mr-1 bg-green-400"></span>
                تایید شده
              </span>
            </div>
          </div>
          
          <!-- پشتیبانی -->
          <div class="bg-gray-50 rounded-lg p-3 border border-gray-200">
            <div class="text-3xl mb-2">💬</div>
            <div class="footer-widget-text font-medium text-sm mb-1">پشتیبانی 24/7</div>
            <div class="footer-widget-text text-opacity-70 text-xs">همیشه در دسترس</div>
            <div class="mt-2">
              <span class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-green-500 bg-opacity-20 text-green-300">
                <span class="w-2 h-2 rounded-full mr-1 bg-green-400"></span>
                فعال
              </span>
            </div>
          </div>
        </div>
      </div>
      
      <!-- گواهینامه‌ها -->
      <div v-if="certificates && certificates.length > 0" class="mt-6 pt-4 border-t border-gray-200">
        <h5 class="footer-widget-text font-medium mb-3 text-sm">گواهینامه‌های ما</h5>
        <div class="flex flex-wrap justify-center gap-3">
          <div
            v-for="cert in certificates"
            :key="cert.id"
            class="bg-gray-50 rounded-lg p-2 border border-gray-200 hover:bg-opacity-10 transition-all duration-200 cursor-pointer"
            @click="handleCertificateClick(cert)"
          >
            <div class="text-lg mb-1">{{ cert.icon }}</div>
            <div class="footer-widget-text text-xs">{{ cert.name }}</div>
            <div v-if="cert.issueDate" class="footer-widget-text text-opacity-60 text-xs">
              {{ formatDate(cert.issueDate) }}
            </div>
          </div>
        </div>
      </div>
      
      <!-- آمار اعتماد -->
      <div v-if="showTrustStats && (customerCount || satisfactionRate || yearsOfService)" class="mt-6 pt-4 border-t border-gray-200">
        <h5 class="footer-widget-text font-medium mb-3 text-sm">آمار اعتماد</h5>
        <div class="grid grid-cols-3 gap-6 footer-widget-text text-opacity-80">
          <div v-if="customerCount" class="text-center">
            <div class="text-lg font-semibold">{{ customerCount.toLocaleString() }}</div>
            <div class="text-xs">مشتری راضی</div>
          </div>
          <div v-if="satisfactionRate" class="text-center">
            <div class="text-lg font-semibold">{{ satisfactionRate }}%</div>
            <div class="text-xs">رضایت مشتری</div>
          </div>
          <div v-if="yearsOfService" class="text-center">
            <div class="text-lg font-semibold">{{ yearsOfService }}</div>
            <div class="text-xs">سال خدمت</div>
          </div>
        </div>
      </div>
      
      <!-- دکمه اطلاعات بیشتر -->
      <div v-if="moreInfoUrl" class="mt-6">
        <a
          :href="moreInfoUrl"
          class="inline-flex items-center px-4 py-2 bg-gray-100 hover:bg-opacity-30 footer-widget-text text-sm rounded-lg transition-all duration-200 border border-gray-300"
        >
          <span>ℹ️</span>
          <span class="mr-2">اطلاعات بیشتر</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface TrustBadge {
  id: string
  title: string
  description?: string
  icon?: string
  htmlCode?: string  // کد HTML نماد (مثل اینماد، ساماندهی و...)
  status?: 'verified' | 'pending' | 'expired'
}

interface Certificate {
  id: string
  name: string
  icon: string
  issueDate?: string
}

interface Props {
  title?: string
  description?: string
  trustBadges?: TrustBadge[]
  certificates?: Certificate[]
  showTrustStats?: boolean
  customerCount?: number
  satisfactionRate?: number
  yearsOfService?: number
  moreInfoUrl?: string
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: 'نشان‌های اعتماد',
  description: 'ما متعهد به ارائه خدمات با کیفیت و امن به مشتریان خود هستیم',
  trustBadges: undefined,
  certificates: undefined,
  showTrustStats: false,
  customerCount: 0,
  satisfactionRate: 0,
  yearsOfService: 0,
  moreInfoUrl: '',
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// مدیریت کلیک روی نشان اعتماد
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const handleBadgeClick = (_badge: TrustBadge) => {
  // اینجا می‌توانید عملیات مورد نظر را انجام دهید

}

// مدیریت کلیک روی گواهینامه
const _handleCertificateClick = (_cert: Certificate) => {
  // اینجا می‌توانید عملیات مورد نظر را انجام دهید

}
const handleCertificateClick = _handleCertificateClick

// تبدیل متن وضعیت
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const getStatusText = (status: string): string => {
  switch (status) {
    case 'verified':
      return 'تایید شده'
    case 'pending':
      return 'در انتظار تایید'
    case 'expired':
      return 'منقضی شده'
    default:
      return status
  }
}

// فرمت تاریخ
const formatDate = (dateString: string): string => {
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('fa-IR')
  } catch {
    return dateString
  }
}

// استایل کامپوننت بر اساس چینش
const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-trust-widget {
  transition: all 0.2s ease;
}

.footer-trust-widget .grid > div:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.footer-trust-widget a:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* رنگ متن پیش‌فرض - تیره برای پس‌زمینه روشن */
.footer-widget-text {
  color: #1f2937; /* gray-800 */
}

/* انیمیشن برای نشان‌های اعتماد */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.7;
  }
}

.footer-trust-widget .grid > div:hover .text-3xl {
  animation: pulse 1s ease-in-out;
}
</style>

