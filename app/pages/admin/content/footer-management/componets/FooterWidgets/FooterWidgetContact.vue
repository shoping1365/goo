<template>
  <div class="footer-contact-widget" :style="containerStyle">
    <div class="footer-contact-inner" :style="innerSpacingStyle">
      <h4 v-if="title" class="footer-widget-title font-semibold mb-4 text-lg">{{ title }}</h4>

      <div v-if="locationEntries.length" class="space-y-4 footer-location-list">
        <div
          v-for="location in locationEntries"
          :key="location.id"
          class="footer-location-entry"
        >
          <div class="footer-location-row">
            <span class="text-lg">📍</span>
            <div class="footer-location-content">
              <p v-if="location.title" class="text-base font-semibold mb-1">{{ location.title }}</p>
              <p v-if="location.address" class="text-sm mb-2">{{ location.address }}</p>
              <div v-if="location.phones.length" class="footer-location-phones">
                <div
                  v-for="phone in location.phones"
                  :key="`${location.id}-phone-${phone}`"
                  class="footer-phone-row"
                >
                  <span class="footer-phone-icon">📞</span>
                  <a :href="`tel:${phone}`" class="phone-link">{{ phone }}</a>
                </div>
              </div>
              <div v-if="location.email" class="footer-location-email">
                <a :href="`mailto:${location.email}`" class="hover:opacity-80 transition-colors">
                  {{ location.email }}
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="shouldShowGeneralSection" class="space-y-3 footer-general-contact">
        <div v-if="address" class="flex items-start space-x-2 space-x-reverse footer-widget-text">
          <span class="text-lg">📍</span>
          <p class="text-sm">{{ address }}</p>
        </div>

        <div v-if="secondaryAddress" class="flex items-start space-x-2 space-x-reverse footer-widget-text">
          <span class="text-lg">📍</span>
          <p class="text-sm">{{ secondaryAddress }}</p>
        </div>

        <div v-if="allPhones.length" class="flex items-start space-x-2 space-x-reverse footer-widget-text">
          <span class="text-lg">📞</span>
          <span class="text-sm phone-line">
            <template v-for="(phone, index) in allPhones" :key="`fallback-phone-${index}`">
              <span v-if="index > 0" class="phone-separator">-</span>
              <a :href="`tel:${phone}`" class="hover:opacity-80 transition-colors phone-link">
                {{ phone }}
              </a>
            </template>
          </span>
        </div>

        <div v-if="email" class="flex items-start space-x-2 space-x-reverse footer-widget-text">
          <span class="text-lg">✉️</span>
          <a :href="`mailto:${email}`" class="text-sm hover:opacity-80 transition-colors">
            {{ email }}
          </a>
        </div>
      </div>

      <div v-else class="footer-widget-text opacity-60 text-sm">
        اطلاعات تماس تعریف نشده است
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface LocationItem {
  id?: string | number
  title?: string
  address?: string
  phones?: string[]
  email?: string
}

interface Props {
  title?: string
  address?: string
  secondaryAddress?: string
  phone?: string
  phones?: string[]
  mobiles?: string[]
  email?: string
  workingHours?: string
  locations?: LocationItem[]
  paddingRight?: number
  paddingLeft?: number
  paddingTop?: number
  paddingBottom?: number
  align?: 'left' | 'center' | 'right'
}

const props = withDefaults(defineProps<Props>(), {
  title: '',
  address: '',
  secondaryAddress: '',
  phone: '',
  phones: () => [],
  mobiles: () => [],
  email: '',
  workingHours: '',
  locations: () => [],
  paddingRight: 0,
  paddingLeft: 0,
  paddingTop: 0,
  paddingBottom: 0,
  align: 'center'
})

const normalizeText = (value?: string | null) => typeof value === 'string' ? value.replace(/\s+/g, ' ').trim() : ''
const normalizePhone = (value?: string | null) => typeof value === 'string' ? value.replace(/\s+/g, ' ').trim() : ''

const rawLocationEntries = computed(() => Array.isArray(props.locations) ? props.locations : [])

const locationEntries = computed(() => rawLocationEntries.value
  .map((location, index) => {
    const id = location?.id ?? `loc-${index}`
    const title = normalizeText(location?.title)
    const address = normalizeText(location?.address)
    const email = normalizeText(location?.email)
    const phonesArray = Array.isArray(location?.phones) ? location.phones : []
    const phones = Array.from(new Set(phonesArray.map(phone => normalizePhone(phone)).filter(Boolean)))

    return {
      id,
      title,
      address,
      email,
      phones
    }
  })
  .filter(entry => entry.address || entry.title || entry.phones.length || entry.email)
)

const hasLocationEntries = computed(() => locationEntries.value.length > 0)

const phonesArray = computed(() => (props.phones || []).map(phone => normalizePhone(phone)).filter(Boolean))

const mobilesArray = computed(() => (props.mobiles || []).map(phone => normalizePhone(phone)).filter(Boolean))

const allPhones = computed(() => Array.from(new Set([...phonesArray.value, ...mobilesArray.value])))

const shouldShowGeneralSection = computed(() => {
  if (hasLocationEntries.value) {
    return false
  }
  return Boolean(props.address || props.secondaryAddress || allPhones.value.length > 0 || props.email)
})

const hasContactInfo = computed(() => hasLocationEntries.value || shouldShowGeneralSection.value)

const containerStyle = computed(() => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: getJustifyContent(props.align),
  width: '100%',
  height: '100%'
}))

const innerSpacingStyle = computed(() => ({
  paddingRight: `${props.paddingRight}px`,
  paddingLeft: `${props.paddingLeft}px`,
  paddingTop: `${props.paddingTop}px`,
  paddingBottom: `${props.paddingBottom}px`
}))

function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'
    default:
      return 'center'
  }
}
</script>

<style scoped>
.footer-contact-widget {
  transition: all 0.2s ease;
  color: #000;
}

.footer-contact-widget a:hover {
  transform: scale(1.05);
}

.footer-contact-inner {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  text-align: right;
  width: 100%;
  gap: 12px;
  direction: rtl;
}

.footer-location-list {
  text-align: right;
}

.footer-location-entry {
  display: flex;
  justify-content: flex-start;
  width: 100%;
}

.footer-location-row {
  display: flex;
  gap: 12px;
  align-items: flex-start;
  justify-content: flex-start;
  text-align: right;
  width: 100%;
  direction: rtl;
}

.footer-location-content {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: right;
  gap: 4px;
}

.footer-phone-row {
  display: flex;
  align-items: center;
  gap: 8px;
}

.footer-phone-icon {
  font-size: 0.95rem;
}

.footer-location-phones {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.footer-location-email {
  margin-top: 4px;
}

.footer-location-content a {
  text-decoration: none;
}

.footer-widget-title,
.footer-widget-text {
  color: #000;
}

.footer-widget-text a {
  color: inherit;
  text-decoration: none;
}

.phone-line {
  display: inline-flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 10px;
}

.phone-separator {
  margin: 0;
}

.footer-general-contact {
  text-align: right;
}
</style>
