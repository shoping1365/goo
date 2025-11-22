<template>
   <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<script setup lang="ts">
// Filter out Nuxt HMR router warnings in development
if (process.dev && process.client) {
  const originalWarn = console.warn
  console.warn = (...args: unknown[]) => {
    const message = args[0]
    if (typeof message === 'string' && message.includes('No match found for location with path "/_nuxt/')) {
      return
    }
    originalWarn.apply(console, args)
  }
}
</script>

<style>

body {
    min-height: 100vh;
    margin: 0 !important;
    padding: 0 !important;
}

* {
    box-sizing: border-box;
}

/* Force chat widget positioning for all widgets */
.user-chat-widget,
[class*="chat-widget"] {
    position: fixed;
    bottom: 24px;
    right: 24px;
    left: auto;
    top: auto;
    direction: ltr;
    transform: none;
    margin: 0;
    padding: 0;
}

/* Component styles */
.collapsible-section {
    border: 1px solid rgb(229 229 229);
    border-radius: 0.375rem;
    overflow: hidden;
}

/* کلاس‌های آیتم‌های هدر */
.logo-icon-wrapper,
.menu-icon-wrapper,
.cart-icon-wrapper,
.auth-icon-wrapper,
.notification-icon-wrapper,
.language-icon-wrapper,
.currency-icon-wrapper,
.social-icon-wrapper,
.custom-icon-wrapper {
    background-color: rgb(249 250 251);
    padding: 0.5rem;
    border-radius: 0.5rem;
    transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
    transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    transition-duration: 150ms;
}

.logo-icon-wrapper:hover,
.menu-icon-wrapper:hover,
.cart-icon-wrapper:hover,
.auth-icon-wrapper:hover,
.notification-icon-wrapper:hover,
.language-icon-wrapper:hover,
.currency-icon-wrapper:hover,
.social-icon-wrapper:hover,
.custom-icon-wrapper:hover {
    background-color: rgb(239 246 255);
}

.cart-badge,
.notification-badge {
    position: absolute;
    top: -0.25rem;
    right: -0.25rem;
    width: 1.25rem;
    height: 1.25rem;
    background-color: rgb(239 68 68);
    color: white;
    font-size: 0.75rem;
    line-height: 1rem;
    border-radius: 9999px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 500;
}
</style>
