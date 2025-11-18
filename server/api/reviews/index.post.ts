import type { H3Event } from 'h3'
import { defineEventHandler } from 'h3'
import { getRequestHeader, readMultipartFormData } from 'h3'
import { useRuntimeConfig } from '#imports'

declare const requireAuth: (event: H3Event) => Promise<{ id: number | string; role: string; token?: string } | null>

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  // احراز هویت سمت سرور نیترو (نیاز به لاگین)
  const user = await requireAuth(event)
  const authHeader = user?.token ? { Authorization: `Bearer ${user.token}` } : {}

  const cookie = getRequestHeader(event, 'cookie') || ''

  // عبور مستقیم فرم-داده به Go
  const form = await readMultipartFormData(event)
  const formData = new FormData()
  for (const part of form || []) {
    if (part.type === 'file' && part.filename && part.data) {
      const buffer = Buffer.isBuffer(part.data) ? part.data : Buffer.from(part.data as ArrayBuffer)
      const uint8Array = new Uint8Array(buffer)
      formData.append(part.name || 'file', new Blob([uint8Array], { type: part.type || 'application/octet-stream' }), part.filename)
    } else if (part.name) {
      formData.append(part.name, part.data?.toString() || '')
    }
  }

  const res = await fetch(`${base}/api/reviews`, {
    method: 'POST',
    headers: {
      cookie,
      ...authHeader,
      'X-User-ID': String(user?.id || 1),
      'X-User-Role': user?.role || 'customer'
    },
    body: formData as unknown as BodyInit
  })

  const text = await res.text()
  try {
    return JSON.parse(text)
  } catch {
    return { error: 'bad_response', raw: text }
  }
})
