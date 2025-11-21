<template>
  <div class="footer-widget footer-widget--trust" :style="containerStyle">
    <div v-if="viewModel.title || viewModel.description" class="footer-widget__header">
      <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
      <p v-if="viewModel.description" class="footer-widget__description">{{ viewModel.description }}</p>
    </div>

    <div
      v-if="viewModel.badges.length"
      class="trust-slider"
      :class="{ 'trust-slider--single': viewModel.badges.length === 1 }"
      :style="sliderStyle"
      @mouseenter="pauseAutoplay"
      @mouseleave="resumeAutoplay"
    >
      <button
        v-if="viewModel.showArrows && viewModel.badges.length > 1"
        type="button"
        class="trust-slider__control trust-slider__control--prev"
        aria-label="نمایش لوگوی قبلی"
        @click="showPrevious()"
      >
        ‹
      </button>

      <transition name="trust-fade" mode="out-in">
        <div v-if="activeBadge" :key="activeBadge.id" class="trust-slide">
          <!-- 
            ⚠️ امنیت XSS: استفاده از v-html خطرناک است!
            
            این کد محتوای HTML را بدون sanitization نمایش می‌دهد که می‌تواند منجر به حملات XSS شود.
            
            ✅ راه حل صحیح:
            1. قبل از استفاده از v-html، محتوا را با کتابخانه sanitization (مثل DOMPurify) پاکسازی کنید
            2. یا از {{ }} به جای v-html استفاده کنید اگر HTML نیاز نیست
            3. محتوای کاربر را هرگز بدون sanitization در v-html قرار ندهید
            
            مثال صحیح:
            import DOMPurify from 'dompurify'
            const sanitizedHtml = computed(() => DOMPurify.sanitize(activeBadge.html))
            <div v-html="sanitizedHtml"></div>
          -->
          <SanitizedHtml v-if="activeBadge.html" :content="activeBadge.html" class="trust-slide__html" />

          <a
            v-else-if="activeBadge.link"
            :href="activeBadge.link"
            class="trust-slide__link"
            target="_blank"
            rel="noopener noreferrer"
          >
            <img
              v-if="activeBadge.imageUrl"
              :src="activeBadge.imageUrl"
              :alt="activeBadge.title || 'trust badge'"
              class="trust-slide__image"
              loading="lazy"
            />
            <span v-if="activeBadge.title" class="trust-slide__title">{{ activeBadge.title }}</span>
          </a>

          <div v-else class="trust-slide__plain">
            <img
              v-if="activeBadge.imageUrl"
              :src="activeBadge.imageUrl"
              :alt="activeBadge.title || 'trust badge'"
              class="trust-slide__image"
              loading="lazy"
            />
            <span v-if="activeBadge.title" class="trust-slide__title">{{ activeBadge.title }}</span>
          </div>
        </div>
      </transition>

      <button
        v-if="viewModel.showArrows && viewModel.badges.length > 1"
        type="button"
        class="trust-slider__control trust-slider__control--next"
        aria-label="نمایش لوگوی بعدی"
        @click="showNext()"
      >
        ›
      </button>
    </div>

    <div v-else class="footer-widget__empty">{{ viewModel.emptyText }}</div>

    <div
      v-if="viewModel.showIndicators && viewModel.badges.length > 1"
      class="trust-indicators"
      role="tablist"
      aria-label="شاخص‌های اسلایدر نشان اعتماد"
    >
      <button
        v-for="(badge, index) in viewModel.badges"
        :key="badge.id"
        type="button"
        class="trust-indicator"
        :class="{ 'trust-indicator--active': index === currentIndex }"
        :aria-label="`نمایش نشان شماره ${index + 1}`"
        :aria-current="index === currentIndex"
        @click="goTo(index)"
      ></button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import SanitizedHtml from '~/components/common/SanitizedHtml.vue'

type RawBadge = {
  id?: string | number
  title?: string
  link?: string
  imageUrl?: string
  html?: string
}

interface Badge {
  id: string | number
  title: string
  link: string
  imageUrl: string
  html: string
}

const props = withDefaults(defineProps<{
  title?: string
  description?: string
  badges?: RawBadge[]
  emptyText?: string
  paddingRight?: number
  paddingLeft?: number
  align?: 'left' | 'center' | 'right'
  autoPlay?: boolean
  autoPlayInterval?: number
  showIndicators?: boolean
  showArrows?: boolean
}>(), {
  title: 'نشان‌های اعتماد',
  description: '',
  badges: () => [],
  emptyText: 'نشان اعتمادی ثبت نشده است.',
  paddingRight: 0,
  paddingLeft: 0,
  align: 'center',
  autoPlay: true,
  autoPlayInterval: 5000,
  showIndicators: true,
  showArrows: true
})

const sanitizeBadges = (badges: RawBadge[]): Badge[] => {
  if (!Array.isArray(badges)) return []
  return badges
    .map((badge, index) => {
      const html = badge.html?.trim() || ''
      
      return {
        id: badge.id ?? `badge-${index}`,
        title: badge.title?.trim() || '',
        link: badge.link?.trim() || '',
        imageUrl: badge.imageUrl?.trim() || '',
        html // Sanitization will be done by SanitizedHtml component
      }
    })
    .filter(badge => badge.title || badge.imageUrl || badge.html || badge.link) as Badge[]
}

const normalizeInterval = (value?: number) => {
  const interval = Number(value)
  if (!Number.isFinite(interval) || interval <= 0) return 5000
  return Math.max(2000, Math.round(interval))
}

const viewModel = computed(() => ({
  title: props.title?.trim() || '',
  description: props.description?.trim() || '',
  badges: sanitizeBadges(props.badges || []),
  emptyText: props.emptyText || 'نشان اعتمادی ثبت نشده است.',
  autoPlay: Boolean(props.autoPlay),
  autoPlayInterval: normalizeInterval(props.autoPlayInterval),
  showIndicators: Boolean(props.showIndicators),
  showArrows: Boolean(props.showArrows)
}))

const containerStyle = computed(() => {
  const alignKey = props.align || 'center'
  const alignItems: 'flex-start' | 'center' | 'flex-end' = alignKey === 'left'
    ? 'flex-start'
    : alignKey === 'right'
      ? 'flex-end'
      : 'center'
  const textAlign: 'left' | 'center' | 'right' = alignKey === 'left'
    ? 'left'
    : alignKey === 'right'
      ? 'right'
      : 'center'

  return {
    paddingRight: `${props.paddingRight}px`,
    paddingLeft: `${props.paddingLeft}px`,
    alignItems,
    textAlign
  }
})

const sliderStyle = computed(() => {
  const alignKey = props.align || 'center'
  if (alignKey === 'left') {
    return {
      width: '100%',
      maxWidth: '360px',
      marginInlineStart: '0',
      marginInlineEnd: 'auto'
    }
  }
  if (alignKey === 'right') {
    return {
      width: '100%',
      maxWidth: '360px',
      marginInlineStart: 'auto',
      marginInlineEnd: '0'
    }
  }
  return {
    width: '100%',
    maxWidth: '360px',
    marginInline: '0 auto'
  }
})

const currentIndex = ref(0)
const activeBadge = computed(() => viewModel.value.badges[currentIndex.value] || null)

let autoplayTimer: number | null = null

const clearAutoplay = () => {
  if (autoplayTimer !== null && typeof window !== 'undefined') {
    window.clearInterval(autoplayTimer)
  }
  autoplayTimer = null
}

const startAutoplay = () => {
  if (!import.meta.client) return
  clearAutoplay()
  if (!viewModel.value.autoPlay || viewModel.value.badges.length <= 1) return
  autoplayTimer = window.setInterval(() => {
    showNext(false)
  }, viewModel.value.autoPlayInterval)
}

const restartAutoplay = () => {
  clearAutoplay()
  startAutoplay()
}

const pauseAutoplay = () => {
  if (!import.meta.client) return
  clearAutoplay()
}

const resumeAutoplay = () => {
  if (!import.meta.client) return
  if (viewModel.value.autoPlay) {
    startAutoplay()
  }
}

const showNext = (shouldRestart = true) => {
  const total = viewModel.value.badges.length
  if (total <= 1) return
  currentIndex.value = (currentIndex.value + 1) % total
  if (shouldRestart) restartAutoplay()
}

const showPrevious = () => {
  const total = viewModel.value.badges.length
  if (total <= 1) return
  currentIndex.value = (currentIndex.value - 1 + total) % total
  restartAutoplay()
}

const goTo = (index: number) => {
  if (index < 0 || index >= viewModel.value.badges.length) return
  currentIndex.value = index
  restartAutoplay()
}

watch(() => viewModel.value.badges, (badges) => {
  if (!badges.length) {
    currentIndex.value = 0
    clearAutoplay()
    return
  }
  if (currentIndex.value >= badges.length) {
    currentIndex.value = 0
  }
  restartAutoplay()
}, { deep: true })

watch(() => [viewModel.value.autoPlay, viewModel.value.autoPlayInterval], () => {
  if (!viewModel.value.autoPlay) {
    clearAutoplay()
  } else {
    restartAutoplay()
  }
})

onMounted(() => {
  startAutoplay()
})

onBeforeUnmount(() => {
  clearAutoplay()
})
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
  width: 100%;
}

.footer-widget__header {
  display: flex;
  flex-direction: column;
  gap: 0.35rem;
}

.footer-widget__title {
  margin: 0;
  font-size: 1rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__description {
  margin: 0;
  font-size: 0.85rem;
  color: var(--footer-widget-muted-color, #6b7280);
}

.trust-slider {
  position: relative;
  width: 100%;
  min-height: 120px;
  padding: 20px 56px;
  border-radius: 16px;
  background: rgba(15, 23, 42, 0.12);
  backdrop-filter: blur(6px);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.trust-slider--single {
  padding-inline: 32px;
}

.trust-slider__control {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 36px;
  height: 36px;
  border-radius: 999px;
  border: none;
  background: rgba(255, 255, 255, 0.7);
  color: #1f2937;
  font-size: 24px;
  line-height: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background 0.2s ease, transform 0.2s ease;
}

.trust-slider__control:hover {
  background: rgba(255, 255, 255, 0.9);
  transform: translateY(-50%) scale(1.05);
}

.trust-slider__control:focus-visible {
  outline: 2px solid #2563eb;
  outline-offset: 2px;
}

.trust-slider__control--prev {
  right: auto;
  left: 12px;
}

.trust-slider__control--next {
  left: auto;
  right: 12px;
}

.trust-slide {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
}

.trust-slide__link,
.trust-slide__plain {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  color: inherit;
  text-decoration: none;
  align-items: center;
}

.trust-slide__image {
  max-width: 140px;
  max-height: 120px;
  object-fit: contain;
  filter: drop-shadow(0 4px 18px rgba(15, 23, 42, 0.18));
}

.trust-slide__title {
  font-size: 0.85rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.trust-slide__html :deep(img),
.trust-slide__html :deep(svg),
.trust-slide__html :deep(canvas) {
  max-width: 160px;
  max-height: 130px;
  object-fit: contain;
  display: block;
  margin: 0 auto;
}

.trust-slide__html {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.trust-indicators {
  display: flex;
  gap: 8px;
  justify-content: center;
  margin-top: 0.5rem;
}

.trust-indicator {
  width: 10px;
  height: 10px;
  border-radius: 999px;
  background: rgba(37, 99, 235, 0.25);
  border: none;
  cursor: pointer;
  transition: background 0.2s ease, transform 0.2s ease;
}

.trust-indicator:hover {
  background: rgba(37, 99, 235, 0.45);
}

.trust-indicator--active {
  background: #2563eb;
  transform: scale(1.1);
}

.trust-fade-enter-active,
.trust-fade-leave-active {
  transition: opacity 0.25s ease, transform 0.25s ease;
}

.trust-fade-enter-from,
.trust-fade-leave-to {
  opacity: 0;
  transform: translateY(8px);
}

.footer-widget__empty {
  text-align: center;
  font-size: 0.85rem;
  color: var(--footer-widget-muted-color, #9ca3af);
}

@media (max-width: 640px) {
  .trust-slider {
    padding: 18px 48px;
  }

  .trust-slider__control {
    width: 32px;
    height: 32px;
    font-size: 20px;
  }

  .trust-slide__image {
    max-width: 120px;
    max-height: 100px;
  }
}
</style>
