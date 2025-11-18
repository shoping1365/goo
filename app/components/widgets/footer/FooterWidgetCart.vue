<template>
  <NuxtLink :to="viewModel.cartUrl" class="footer-widget footer-widget--cart">
    <span class="footer-widget__icon" aria-hidden="true">Cart</span>
    <div class="footer-widget__details">
      <span class="footer-widget__label">{{ viewModel.label }}</span>
      <span class="footer-widget__meta">{{ viewModel.meta }}</span>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  cartUrl?: string
  itemCount?: number
  totalLabel?: string
  currency?: string
  label?: string
}>(), {
  cartUrl: '/cart',
  itemCount: 0,
  totalLabel: '',
  currency: '',
  label: 'Your cart'
})

const viewModel = computed(() => {
  const count = Number.isFinite(props.itemCount) ? Math.max(0, props.itemCount || 0) : 0
  const totalLabel = props.totalLabel?.trim() || ''
  const currency = props.currency?.trim() || ''
  const hasTotal = Boolean(totalLabel)
  const meta = hasTotal
    ? `${count} items · ${totalLabel}${currency ? ` ${currency}` : ''}`
    : `${count} items`

  return {
    cartUrl: props.cartUrl || '/cart',
    label: props.label || 'Your cart',
    meta
  }
})
</script>

<style scoped>
.footer-widget {
  display: inline-flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border-radius: 0.75rem;
  background: var(--footer-widget-surface, #f3f4f6);
  color: var(--footer-widget-text-color, #111827);
  text-decoration: none;
}

.footer-widget__icon {
  font-size: 1.5rem;
}

.footer-widget__details {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.footer-widget__label {
  font-weight: 600;
  font-size: 0.95rem;
}

.footer-widget__meta {
  font-size: 0.8rem;
  color: var(--footer-widget-muted-color, #6b7280);
}
</style>
