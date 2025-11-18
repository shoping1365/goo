import { eventHandler, type H3Event } from 'h3'
import { useRuntimeConfig } from '#imports'

export default eventHandler((event: H3Event) => {
  // موقتاً غیرفعال برای رفع مشکل encoding
  return
  const config = useRuntimeConfig()
  const goApiBase = config.public.goApiBase
  const res = event.node.res
  res.setHeader('X-Frame-Options', 'DENY')
  res.setHeader('X-Content-Type-Options', 'nosniff')
  res.setHeader('Referrer-Policy', 'no-referrer-when-downgrade')
  res.setHeader('Permissions-Policy', 'geolocation=(), camera=(), microphone=()')
  res.setHeader('X-XSS-Protection', '1; mode=block')
  // Conservative CSP suitable for server APIs; relax per need on pages
  res.setHeader('Content-Security-Policy', `default-src 'self'; img-src 'self' data:; media-src 'self' data:; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'; connect-src 'self' ${goApiBase} http://localhost:3000`)
})