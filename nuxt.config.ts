import { fileURLToPath, URL } from 'node:url'
import { defineNuxtConfig } from 'nuxt/config'
// @ts-ignore - Tailwind v4 Vite plugin
import tailwindcss from '@tailwindcss/vite'

const resolveApiBase = () => {
  const raw = process.env.NUXT_PUBLIC_GO_API_BASE || process.env.NUXT_PUBLIC_API_BASE
  if (!raw) return ''
  const trimmed = raw.trim()
  if (!trimmed) return ''
  const lowered = trimmed.toLowerCase()
  if (['undefined', 'null', 'false'].includes(lowered)) {
    return ''
  }
  return trimmed
}

const resolvedPublicApiBase = resolveApiBase()

export default defineNuxtConfig({
  compatibilityDate: '2025-10-13',
  devtools: { enabled: true },
  telemetry: false,
  ssr: true,

  // Route rules استاندارد Nuxt برای static assets
  routeRules: {
    // Proxy endpoint‌ها به Go backend
    '/api/products/**': {
      proxy: { to: `${process.env.NUXT_PUBLIC_GO_API_BASE}/api/products/**` }
    },
    '/api/categories/**': {
      proxy: { to: `${process.env.NUXT_PUBLIC_GO_API_BASE}/api/categories/**` }
    },
    '/api/auth/**': {
      proxy: { to: `${process.env.NUXT_PUBLIC_GO_API_BASE}/api/auth/**` }
    },
    '/api/users/**': {
      proxy: { to: `${process.env.NUXT_PUBLIC_GO_API_BASE}/api/users/**` }
    },
    '/api/admin/digikala/**': {
      proxy: { to: `${process.env.NUXT_PUBLIC_GO_API_BASE}/api/admin/digikala/**` }
    },

    // قانون کش برای فایل‌های آپلود شده - Caddy خودش proxy میکنه
    '/uploads/**': {
      headers: {
        'Cache-Control': 'public, max-age=604800, must-revalidate'
      }
    },

    // قانون کش برای فایل‌های build شده Nuxt
    '/_nuxt/**': {
      headers: {
        'Cache-Control': 'public, max-age=31536000, immutable'
      }
    }
  },



  // غیرفعال کردن app manifest و features اضافی
  experimental: {
    appManifest: false,
    buildCache: false,
    componentIslands: false
  },

  // جلوگیری از manifest و payload extraction در dev
  features: {
    devLogs: false
  },



  typescript: {
    typeCheck: true,
    strict: false
  },

  app: {
    head: {
      title: 'فروشگاه ',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' }
      ]
    },
    // غیرفعال کردن route rules برای جلوگیری از fetch manifest
    pageTransition: false,
    layoutTransition: false
  },

  runtimeConfig: {
    public: {
      goApiBase: resolvedPublicApiBase,
      apiBase: resolvedPublicApiBase,
      wsBase: process.env.NUXT_PUBLIC_WS_BASE,
      siteUrl: process.env.NUXT_PUBLIC_SITE_URL,
      // Security features are now enabled
      disableSecurity: false
    }
  },

  // @ts-expect-error - security config is from nuxt-security module
  security: {
    headers: {
      contentSecurityPolicy: {
        'img-src': ["'self'", 'data:', 'https:'],
        'script-src': ["'self'", "'unsafe-inline'", "'unsafe-eval'"],
        'style-src': ["'self'", "'unsafe-inline'"],
        'font-src': ["'self'", 'data:', 'https:'],
        'connect-src': ["'self'", 'ws:', 'wss:', 'https:'],
        'frame-src': ["'self'"],
        'object-src': ["'none'"],
        'base-uri': ["'self'"]
      },
      crossOriginEmbedderPolicy: false,
      crossOriginOpenerPolicy: false,
      crossOriginResourcePolicy: false,
      xFrameOptions: 'SAMEORIGIN',
      xContentTypeOptions: 'nosniff',
      xXSSProtection: '1; mode=block',
      referrerPolicy: 'strict-origin-when-cross-origin',
      permissionsPolicy: {
        camera: [],
        microphone: [],
        geolocation: []
      }
    },
    rateLimiter: {
      tokensPerInterval: 150,
      interval: 'minute',
      throwError: false
    },
    corsHandler: {
      origin: process.env.NUXT_PUBLIC_SITE_URL || '*',
      methods: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH'],
      credentials: true
    },
    requestSizeLimiter: {
      maxRequestSizeInBytes: 10000000,
      maxUploadFileRequestInBytes: 50000000,
      throwError: false
    },
    xssValidator: {
      throwError: false
    },
    hidePoweredBy: true
  },

  // Nitro Configuration  
  nitro: {
    // مهم: تنظیم preset برای SSR production
    preset: 'node-server',

    publicAssets: [
      {
        dir: 'public',
        maxAge: 60 * 60 * 24 * 365 // کش کردن برای یک سال
      }
    ],


    experimental: {
      asyncContext: true
    },
    // تنظیمات برای serve کردن فایل‌های static
    serveStatic: true,
    // اضافه کردن تنظیمات برای رفع مشکل encoding
    compressPublicAssets: false,
    // Cookie Configuration و Storage
    storage: {
      'auth': {
        driver: 'memory'
      },
      'uploads': {
        driver: 'fs',
        base: './public/uploads'
      }
    }
  },

  alias: {
    '~/server': fileURLToPath(new URL('./server', import.meta.url)),
    '~~/assets': fileURLToPath(new URL('./assets', import.meta.url)),
    '~~/types': fileURLToPath(new URL('./types', import.meta.url))
  },

  modules: [
    '@pinia/nuxt',
    '@nuxtjs/color-mode',
    'nuxt-swiper',
    'nuxt-security'
  ],

  postcss: {
    plugins: {
      'autoprefixer': {},
    },
  },

  components: [
    {
      path: '~/components',
      pathPrefix: false,
      global: true,
      priority: 10
    },
    {
      path: '~/components/admin/common',
      pathPrefix: false,
      global: false,
      priority: 1
    },
    {
      path: '~/app/pages/checkout/components',
      pathPrefix: false,
      global: true,
      priority: 5
    }
  ],

  imports: {
    dirs: ['composables/**', 'utils/**']
  },

  css: [
    '~~/assets/css/main.css'
  ],

  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: false,
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: false,
      __VUE_OPTIONS_API__: true
    },
    ssr: {
      noExternal: ['persian-date']
    },
    optimizeDeps: {
      include: ['vue', 'pinia', 'axios', 'persian-date']
    },
    build: {
      sourcemap: false, // غیرفعال کردن sourcemap برای رفع warning های Tailwind
      rollupOptions: {
        external: [],
        output: {
          manualChunks: undefined
        }
      }
    },
    plugins: [
      tailwindcss()
    ]
  }
})