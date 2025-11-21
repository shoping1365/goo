<template>
  <div class="footer-copyright-widget" :style="widgetStyle">
    <div class="text-center">
      <!-- متن کپی‌رایت -->
      <div class="text-white text-opacity-80 text-sm">
        <!-- 
          ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
          
          این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
          
          ✅ راه حل صحیح:
          1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
          2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
          3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
          
          مثال صحیح:
          import DOMPurify from 'dompurify'
          const sanitizedCopyrightHtml = computed(() => DOMPurify.sanitize(copyrightHtml.value))
          <div v-html="sanitizedCopyrightHtml"></div>
        -->
        <!-- eslint-disable-next-line vue/no-v-html -->
        <div class="copyright-html" v-html="copyrightHtml"></div>
        
        <!-- لینک‌های اضافی -->
        <div v-if="additionalLinks && additionalLinks.length > 0" class="mt-2 space-x-4 space-x-reverse">
          <NuxtLink
            v-for="link in additionalLinks"
            :key="link.text"
            :to="link.url"
            class="text-white text-opacity-60 hover:text-opacity-100 transition-opacity text-xs"
          >
            {{ link.text }}
          </NuxtLink>
        </div>
        
        <!-- متن اضافی -->
        <p v-if="additionalText" class="mt-2 text-xs text-white text-opacity-60">
          {{ additionalText }}
        </p>
        
        <!-- آدرس و اطلاعات اضافی -->
        <div v-if="showAddress" class="mt-3 pt-3 border-t border-white border-opacity-20">
          <p class="text-xs text-white text-opacity-60">
            {{ address || 'آدرس شرکت' }}
          </p>
          <p v-if="phone" class="text-xs text-white text-opacity-60 mt-1">
            تلفن: {{ phone }}
          </p>
          <p v-if="email" class="text-xs text-white text-opacity-60 mt-1">
            ایمیل: {{ email }}
          </p>
        </div>
        
        <!-- لینک‌های شبکه‌های اجتماعی -->
        <div v-if="showSocialLinks && socialLinks && socialLinks.length > 0" class="mt-3">
          <div class="flex justify-center space-x-3 space-x-reverse">
            <a
              v-for="social in socialLinks"
              :key="social.platform"
              :href="social.url"
              target="_blank"
              rel="noopener noreferrer"
              :title="social.platform"
              class="text-white text-opacity-60 hover:text-opacity-100 transition-opacity"
            >
              <span class="text-lg">{{ getSocialIcon(social.platform) }}</span>
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface AdditionalLink {
  text: string
  url: string
}

interface SocialLink {
  platform: string
  url: string
}

interface Props {
  text?: string
  companyName?: string
  copyrightText?: string
  additionalText?: string
  additionalLinks?: AdditionalLink[]
  showAddress?: boolean
  address?: string
  phone?: string
  email?: string
  showSocialLinks?: boolean
  socialLinks?: SocialLink[]
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  text: '',
  companyName: 'نام شرکت',
  copyrightText: 'تمامی حقوق محفوظ است.',
  additionalText: '',
  additionalLinks: () => [
  ],
  showAddress: false,
  address: '',
  phone: '',
  email: '',
  showSocialLinks: false,
  socialLinks: () => [
    { platform: 'instagram', url: 'https://instagram.com' },
    { platform: 'telegram', url: 'https://t.me' },
    { platform: 'twitter', url: 'https://twitter.com' }
  ],
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

// سال جاری
const currentYear = computed(() => {
  return new Date().getFullYear()
})

const copyrightHtml = computed(() => {
  if (props.text && props.text.trim()) {
    return props.text
  }
  const company = props.companyName?.trim() || 'نام شرکت'
  const message = props.copyrightText?.trim() || 'تمامی حقوق محفوظ است.'
  return `© ${currentYear.value} ${company}. ${message}`
})

// آیکون شبکه‌های اجتماعی
const getSocialIcon = (platform: string): string => {
  const icons: { [key: string]: string } = {
    instagram: '📷',
    telegram: '📱',
    twitter: '🐦',
    facebook: '📘',
    linkedin: '💼',
    youtube: '📺',
    whatsapp: '💬'
  }
  return icons[platform] || '🔗'
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
.footer-copyright-widget {
  transition: all 0.2s ease;
}

.copyright-html :deep(p) {
  margin: 0;
}

.footer-copyright-widget a:hover {
  transform: translateY(-1px);
}

.footer-copyright-widget .border-t {
  border-top-width: 1px;
}
</style>
