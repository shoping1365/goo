<template>
  <div class="footer-social-widget" :style="widgetStyle">
    <div class="footer-social-widget__inner">
      <h4 v-if="titleToShow" class="footer-social-widget__title">{{ titleToShow }}</h4>
      <p v-if="descriptionToShow" class="footer-social-widget__description">{{ descriptionToShow }}</p>

      <div v-if="hasSocialLinks" class="footer-social-widget__icons" :class="alignmentClass">
        <a
          v-for="social in normalizedSocials"
          :key="social.id"
          class="footer-social-widget__icon"
          :class="`is-${social.platform}`"
          :href="social.href"
          :target="social.openInNewTab ? '_blank' : '_self'"
          :rel="social.openInNewTab ? 'noopener noreferrer' : undefined"
          :title="social.title"
        >
          <template v-if="social.platform === 'instagram'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12s.014 3.667.073 4.948c.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24s3.667-.014 4.948-.072c4.354-.2 6.782-2.618 6.979-6.98.059-1.281.073-1.689.073-4.948s-.014-3.667-.073-4.948c-.197-4.354-2.618-6.78-6.979-6.98C15.667.014 15.259 0 12 0zm0 5.838A6.162 6.162 0 1 0 18.162 12 6.161 6.161 0 0 0 12 5.838zm0 10.162A4 4 0 1 1 16 12a4 4 0 0 1-4 4zm6.406-11.845a1.44 1.44 0 1 0 1.439 1.44 1.44 1.44 0 0 0-1.439-1.44z"/></svg>
          </template>
          <template v-else-if="social.platform === 'telegram'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12 12-5.373 12-12S18.627 0 12 0zm5.002 7.482l-1.53 7.214c-.115.518-.421.643-.853.399l-2.356-1.736-1.137 1.094c-.126.126-.232.232-.474.232l.167-2.385 4.338-3.913c.189-.167-.041-.262-.293-.095l-5.362 3.378-2.311-.726c-.502-.158-.513-.502.104-.744l9.027-3.482c.421-.158.79.095.653.743z"/></svg>
          </template>
          <template v-else-if="social.platform === 'twitter' || social.platform === 'x'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M21.543 7.104c.015.213.015.425.015.637 0 6.507-4.954 14.008-14.008 14.008-2.783 0-5.369-.81-7.543-2.213.394.045.773.06 1.182.06 2.31 0 4.43-.773 6.126-2.078-2.165-.045-3.984-1.47-4.613-3.434.303.045.606.075.924.075.455 0 .91-.06 1.335-.167-2.253-.455-3.953-2.435-3.953-4.83v-.06c.66.364 1.41.59 2.214.62-1.32-.88-2.195-2.37-2.195-4.06 0-.89.242-1.71.667-2.43 2.4 2.94 5.973 4.87 10.002 5.08-.075-.364-.12-.743-.12-1.122 0-2.7 2.19-4.89 4.89-4.89 1.408 0 2.678.59 3.57 1.53 1.11-.213 2.16-.62 3.108-1.182-.364 1.14-1.14 2.1-2.16 2.7 1-.12 1.95-.379 2.835-.758-.66.99-1.5 1.86-2.46 2.55z"/></svg>
          </template>
          <template v-else-if="social.platform === 'linkedin'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M20.447 20.452H17.1v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.758V9h3.205v1.561h.045c.447-.846 1.541-1.739 3.173-1.739 3.393 0 4.017 2.235 4.017 5.142v6.488zM5.337 7.433a1.859 1.859 0 11.001-3.718 1.859 1.859 0 01-.001 3.718zM7.119 20.452H3.556V9h3.563v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.225 0z"/></svg>
          </template>
          <template v-else-if="social.platform === 'youtube'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M23.498 6.186a2.974 2.974 0 00-2.097-2.097C19.528 3.5 12 3.5 12 3.5s-7.528 0-9.401.589A2.974 2.974 0 00.502 6.186C0 8.059 0 12 0 12s0 3.941.502 5.814a2.974 2.974 0 002.097 2.097C4.472 20.5 12 20.5 12 20.5s7.528 0 9.401-.589a2.974 2.974 0 002.097-2.097C24 15.941 24 12 24 12s0-3.941-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"/></svg>
          </template>
          <template v-else-if="social.platform === 'facebook'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M22.675 0H1.325C.593 0 0 .593 0 1.326v21.348C0 23.407.593 24 1.325 24h11.482v-9.294H9.847v-3.622h2.96V8.413c0-2.934 1.793-4.533 4.415-4.533 1.255 0 2.334.093 2.648.135v3.07h-1.819c-1.425 0-1.701.677-1.701 1.67v2.188h3.4l-.443 3.622h-2.957V24h5.795C23.407 24 24 23.407 24 22.674V1.326C24 .593 23.407 0 22.675 0z"/></svg>
          </template>
          <template v-else-if="social.platform === 'whatsapp'">
            <svg viewBox="0 0 24 24" aria-hidden="true"><path d="M.057 24l1.687-6.163a11.867 11.867 0 01-1.62-6.003C.122 5.281 5.33 0 12.057 0c3.2 0 6.205 1.246 8.47 3.514a11.82 11.82 0 013.474 8.413c-.003 6.726-5.312 11.934-12.039 11.934a12.3 12.3 0 01-6.002-1.52L.057 24zm6.597-3.807c1.676.995 3.276 1.591 5.399 1.593 5.448 0 9.91-4.458 9.913-9.907.002-5.45-4.451-9.913-9.9-9.915-5.452 0-9.91 4.458-9.913 9.908 0 2.225.651 3.891 1.746 5.634l-.999 3.648 3.754-.961zm11.387-5.464c-.073-.121-.268-.195-.558-.342-.29-.146-1.716-.848-1.981-.944-.265-.098-.458-.146-.651.146-.195.29-.747.944-.916 1.138-.168.195-.338.22-.627.073-.29-.146-1.223-.45-2.33-1.43-.861-.767-1.44-1.713-1.609-2.003-.168-.29-.018-.447.128-.593.132-.132.29-.338.436-.507.146-.168.195-.29.29-.485.098-.195.049-.365-.024-.513-.073-.146-.651-1.572-.892-2.152-.235-.565-.476-.488-.651-.498-.168-.009-.365-.011-.558-.011-.195 0-.51.073-.777.365-.265.29-1.02.997-1.02 2.431 0 1.433 1.045 2.818 1.19 3.013.146.195 2.062 3.155 4.992 4.42.698.301 1.245.48 1.67.614.701.224 1.34.193 1.845.117.563-.084 1.716-.701 1.957-1.378.24-.677.24-1.258.168-1.378z"/></svg>
          </template>
          <template v-else>
            <span class="footer-social-widget__fallback">{{ social.shortLabel }}</span>
          </template>
        </a>
      </div>

      <div v-else class="footer-social-widget__empty">شبکه‌های اجتماعی تعریف نشده است</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type Align = 'left' | 'center' | 'right'

interface SocialInput {
  id?: number | string
  platform?: string
  title?: string
  label?: string
  url?: string
  href?: string
  visible?: boolean
  enabled?: boolean
  openInNewTab?: boolean
}

interface Props {
  title?: string
  description?: string
  socials?: SocialInput[]
  instagram?: string
  telegram?: string
  twitter?: string
  linkedin?: string
  youtube?: string
  facebook?: string
  whatsapp?: string
  paddingRight?: number
  paddingLeft?: number
  align?: Align
}

const props = withDefaults(defineProps<Props>(), {
  title: '',
  description: '',
  socials: () => [],
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center'
})

const titleToShow = computed(() => props.title?.trim() || '')
const descriptionToShow = computed(() => props.description?.trim() || '')

const widgetStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`
}))

const alignmentClass = computed(() => {
  switch (props.align) {
    case 'left':
      return 'is-left'
    case 'right':
      return 'is-right'
    default:
      return 'is-center'
  }
})

const normalizedSocials = computed(() => {

  const fromArray = Array.isArray(props.socials) && props.socials.length
    ? props.socials
    : legacySocials()

  const result = fromArray
    .filter(item => {
      // فقط چک کنیم که enabled نباشد false
      const isVisible = item.visible !== false && item.enabled !== false
      return isVisible
    })
    .map((item, index) => {
      const platform = (item.platform || '').toLowerCase().trim()
      const href = (item.href || item.url || '#').trim()
      const title = item.title || item.label || platform || 'شبکه اجتماعی'
      return {
        id: item.id ?? `${platform || 'social'}-${index}`,
        platform,
        href: href || '#',
        title,
        shortLabel: title.slice(0, 2).toUpperCase(),
        visible: true,
        openInNewTab: item.openInNewTab !== false
      }
    })

  return result
})

const hasSocialLinks = computed(() => normalizedSocials.value.length > 0)

function legacySocials(): SocialInput[] {
  const legacyMap: Array<{ platform: string; value?: string }> = [
    { platform: 'instagram', value: props.instagram },
    { platform: 'telegram', value: props.telegram },
    { platform: 'twitter', value: props.twitter },
    { platform: 'linkedin', value: props.linkedin },
    { platform: 'youtube', value: props.youtube },
    { platform: 'facebook', value: props.facebook },
    { platform: 'whatsapp', value: props.whatsapp }
  ]

  return legacyMap
    .filter(item => typeof item.value === 'string' && item.value.trim().length)
    .map(item => ({
      id: item.platform,
      platform: item.platform,
      url: item.value,
      title: item.platform,
      openInNewTab: true
    }))
}
</script>

<style scoped>
.footer-social-widget {
  width: 100%;
  height: 100%;
}

.footer-social-widget__inner {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  align-items: stretch;
}

.footer-social-widget__title {
  margin: 0;
  color: #000000;
  font-weight: 600;
  font-size: 1rem;
}

.footer-social-widget__description {
  margin: 0;
  color: #4b5563;
  font-size: 0.875rem;
}

.footer-social-widget__icons {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.footer-social-widget__icons.is-left {
  justify-content: flex-start;
}

.footer-social-widget__icons.is-center {
  justify-content: center;
}

.footer-social-widget__icons.is-right {
  justify-content: flex-end;
}

.footer-social-widget__icon {
  width: 40px;
  height: 40px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
  color: #000000;
  border-radius: 9999px;
  transition: transform 0.2s ease, background 0.2s ease, color 0.2s ease;
}

.footer-social-widget__icon svg {
  width: 22px;
  height: 22px;
  fill: currentColor;
}

.footer-social-widget__icon:hover {
  transform: translateY(-2px);
  background: #e5e7eb;
}

.footer-social-widget__icon.is-instagram {
  background: linear-gradient(135deg, #f58529, #dd2a7b, #8134af, #515bd4);
  color: #ffffff;
}

.footer-social-widget__icon.is-telegram {
  background: #29a0da;
  color: #ffffff;
}

.footer-social-widget__icon.is-twitter {
  background: #1d9bf0;
  color: #ffffff;
}

.footer-social-widget__icon.is-linkedin {
  background: #0a66c2;
  color: #ffffff;
}

.footer-social-widget__icon.is-youtube {
  background: #ff0000;
  color: #ffffff;
}

.footer-social-widget__icon.is-facebook {
  background: #1877f2;
  color: #ffffff;
}

.footer-social-widget__icon.is-whatsapp {
  background: #25d366;
  color: #ffffff;
}

.footer-social-widget__fallback {
  font-size: 0.75rem;
  font-weight: 600;
  color: #000000;
}

.footer-social-widget__empty {
  color: #6b7280;
  font-size: 0.85rem;
}
</style>
