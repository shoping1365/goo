<template>
  <div class="footer-widget footer-widget--auth">
    <p class="footer-widget__description">{{ viewModel.description }}</p>
    <div class="footer-widget__actions">
      <NuxtLink :to="viewModel.loginUrl" class="footer-widget__link">{{ viewModel.loginLabel }}</NuxtLink>
      <NuxtLink :to="viewModel.registerUrl" class="footer-widget__link footer-widget__link--primary">
        {{ viewModel.registerLabel }}
      </NuxtLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = withDefaults(defineProps<{
  loginUrl?: string
  registerUrl?: string
  loginLabel?: string
  registerLabel?: string
  description?: string
}>(), {
  loginUrl: '/auth/login',
  registerUrl: '/auth/register',
  loginLabel: 'Sign in',
  registerLabel: 'Create account',
  description: 'Access your account or create a new profile.'
})

const viewModel = computed(() => ({
  loginUrl: props.loginUrl || '/auth/login',
  registerUrl: props.registerUrl || '/auth/register',
  loginLabel: props.loginLabel || 'Sign in',
  registerLabel: props.registerLabel || 'Create account',
  description: props.description || 'Access your account or create a new profile.'
}))
</script>

<style scoped>
.footer-widget {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.footer-widget__description {
  margin: 0;
  font-size: 0.875rem;
  color: var(--footer-widget-text-color, #4b5563);
}

.footer-widget__actions {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.footer-widget__link {
  text-decoration: none;
  color: var(--footer-widget-link-color, #2563eb);
  font-weight: 600;
}

.footer-widget__link:hover {
  text-decoration: underline;
}

.footer-widget__link--primary {
  padding: 0.4rem 0.9rem;
  border-radius: 999px;
  background: var(--footer-widget-primary-bg, #2563eb);
  color: #fff;
}

.footer-widget__link--primary:hover {
  background: var(--footer-widget-primary-bg-hover, #1d4ed8);
}
</style>
