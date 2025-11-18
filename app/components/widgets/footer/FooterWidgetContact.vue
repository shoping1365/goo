<template>
  <div class="footer-widget footer-widget--contact">
    <h4 v-if="viewModel.title" class="footer-widget__title">{{ viewModel.title }}</h4>
    
    <!-- اگر ساختار جدید locations موجود باشد -->
    <div v-if="viewModel.hasLocations" class="footer-widget__locations">
      <div 
        v-for="(location, locIndex) in viewModel.locations" 
        :key="location.id || locIndex"
        class="footer-widget__location"
      >
        <h5 v-if="location.title" class="footer-widget__location-title">{{ location.title }}</h5>
        <ul class="footer-widget__list">
          <li v-if="location.address" class="footer-widget__item">
            <span class="footer-widget__label">آدرس:</span>
            <span class="footer-widget__value">{{ location.address }}</span>
          </li>
          <li v-if="location.phones && location.phones.length" class="footer-widget__item">
            <span class="footer-widget__label">تلفن:</span>
            <span class="footer-widget__value footer-widget__phones">
              <template v-for="(phone, phoneIndex) in location.phones.filter(p => p && p.trim())" :key="`${locIndex}-phone-${phoneIndex}`">
                <span v-if="phoneIndex > 0" class="footer-widget__phones-separator">-</span>
                <a :href="`tel:${phone}`" class="footer-widget__link footer-widget__phone">{{ phone }}</a>
              </template>
            </span>
          </li>
        </ul>
      </div>
    </div>
    
    <!-- Fallback به ساختار قدیمی -->
    <ul v-else class="footer-widget__list">
      <li v-if="viewModel.address" class="footer-widget__item">
        <span class="footer-widget__label">آدرس:</span>
        <span class="footer-widget__value">{{ viewModel.address }}</span>
      </li>
      <li v-if="viewModel.secondaryAddress" class="footer-widget__item">
        <span class="footer-widget__label">آدرس دوم:</span>
        <span class="footer-widget__value">{{ viewModel.secondaryAddress }}</span>
      </li>
      <li v-if="viewModel.allPhones.length" class="footer-widget__item">
        <span class="footer-widget__label">تلفن:</span>
        <span class="footer-widget__value footer-widget__phones">
          <template v-for="(phone, index) in viewModel.allPhones" :key="`phone-${index}`">
            <span v-if="index > 0" class="footer-widget__phones-separator">-</span>
            <a :href="`tel:${phone}`" class="footer-widget__link footer-widget__phone">{{ phone }}</a>
          </template>
        </span>
      </li>
      <li v-if="!viewModel.hasAny" class="footer-widget__item footer-widget__item--empty">
        اطلاعات تماس موجود نیست.
      </li>
    </ul>
    
    <!-- ایمیل مشترک برای هر دو ساختار -->
    <ul v-if="viewModel.email" class="footer-widget__list">
      <li class="footer-widget__item">
        <span class="footer-widget__label">ایمیل:</span>
        <a :href="`mailto:${viewModel.email}`" class="footer-widget__value footer-widget__link">{{ viewModel.email }}</a>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  title?: string
  address?: string
  phone?: string
  phones?: string[]
  mobiles?: string[]
  email?: string
  hours?: string
  secondaryAddress?: string
  locations?: Array<{
    id?: number | string
    title?: string
    address?: string
    phones?: string[]
  }>
}>(), {
  title: 'اطلاعات تماس',
  phones: () => [],
  mobiles: () => [],
  locations: () => []
})

const viewModel = computed(() => {
  // بررسی ساختار جدید locations
  const hasLocations = props.locations && props.locations.length > 0
  const locations = hasLocations ? props.locations : []
  
  // ساختار قدیمی (Fallback)
  const address = props.address?.trim() || ''
  const secondaryAddress = props.secondaryAddress?.trim() || ''
  const phones = (props.phones || []).filter(p => p && p.trim())
  const mobiles = (props.mobiles || []).filter(m => m && m.trim())
  const allPhones = [...phones, ...mobiles]
  const email = props.email?.trim() || ''
  
  return {
    title: props.title?.trim() || '',
    hasLocations,
    locations,
    address,
    secondaryAddress,
    phones,
    mobiles,
    allPhones,
    email,
    hasAny: Boolean(hasLocations || address || secondaryAddress || allPhones.length > 0 || email)
  }
})
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-widget__title {
  margin: 0;
  font-size: 0.95rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #111827);
}

.footer-widget__locations {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.footer-widget__location {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.footer-widget__location-title {
  margin: 0;
  font-size: 0.9rem;
  font-weight: 600;
  color: var(--footer-widget-text-color, #1f2937);
  border-bottom: 1px solid var(--footer-widget-border-color, #e5e7eb);
  padding-bottom: 0.25rem;
}

.footer-widget__list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.footer-widget__item {
  display: flex;
  flex-direction: column;
  gap: 0.15rem;
  font-size: 0.9rem;
  color: var(--footer-widget-text-color, #374151);
}

.footer-widget__label {
  font-weight: 600;
}

.footer-widget__value {
  word-break: break-word;
}

.footer-widget__phones {
  display: inline-flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 10px;
}

.footer-widget__phones-separator {
  margin: 0;
}

.footer-widget__link {
  color: var(--footer-widget-link-color, #2563eb);
  text-decoration: none;
}

.footer-widget__link:hover {
  text-decoration: underline;
}

.footer-widget__item--empty {
  font-style: italic;
  color: var(--footer-widget-muted-color, #9ca3af);
}
</style>
