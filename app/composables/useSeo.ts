export interface SeoOptions {
  title?: string
  description?: string
  /** keywords can be string or array */
  keywords?: string | string[]
  image?: string
  /** overrides index/follow if provided directly */
  robots?: string
  og?: Record<string, string>
  twitter?: Record<string, string>
  index?: boolean
  follow?: boolean
}

interface MetaTag {
  name?: string
  property?: string
  content: string
}

// تعریف useHead برای Nuxt 3
declare const useHead: (head: { title?: string; meta?: MetaTag[] }) => void

export function useSeo(opts: SeoOptions) {
  const keywordsStr = Array.isArray(opts.keywords) ? opts.keywords.join(',') : opts.keywords

  const meta: MetaTag[] = []

  if (opts.description) meta.push({ name: 'description', content: opts.description })
  if (keywordsStr) meta.push({ name: 'keywords', content: keywordsStr })
  if (opts.image) meta.push({ property: 'og:image', content: opts.image })
  // robots directive (index/follow)
  let robotsStr = opts.robots
  if (!robotsStr) {
    const index = opts.index !== false // default is index
    const follow = opts.follow !== false // default is follow
    robotsStr = `${index ? 'index' : 'noindex'},${follow ? 'follow' : 'nofollow'}`
  }
  meta.push({ name: 'robots', content: robotsStr })

  if (opts.og) {
    for (const [k, v] of Object.entries(opts.og)) {
      meta.push({ property: `og:${k}`, content: v })
    }
  }
  if (opts.twitter) {
    for (const [k, v] of Object.entries(opts.twitter)) {
      meta.push({ name: `twitter:${k}`, content: v })
    }
  }

  useHead({
    title: opts.title,
    meta
  })
} 