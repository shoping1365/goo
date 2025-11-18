import { createError, defineEventHandler } from 'h3'

declare const useRuntimeConfig: () => { public: { goApiBase?: string; apiBase?: string } }

type PopularSearchItem = {
  query: string
  count: number
  context?: string
}

type PopularSearchResponse = {
  success: boolean
  data: PopularSearchItem[]
}
const fetchPopularSearches = async (): Promise<PopularSearchResponse> => {
  const config = useRuntimeConfig()
  const apiBase = (config.public as { goApiBase?: string; apiBase?: string }).goApiBase
    || (config.public as { goApiBase?: string; apiBase?: string }).apiBase

  if (!apiBase) {
    throw new Error('apiBase is not configured')
  }

  const endpoint = new URL('/api/search/popular', apiBase)
  const response = await fetch(endpoint, {
    headers: { 'Content-Type': 'application/json' },
  })

  if (!response.ok) {
    throw new Error(`backend responded with status ${response.status}`)
  }

  const payload = (await response.json()) as PopularSearchResponse
  return payload
}

const popularSearchHandler = defineEventHandler(async () => {
  try {
    return await fetchPopularSearches()
  } catch (error: unknown) {
    console.error('Error fetching popular searches:', error)
    throw createError({
      statusCode: 502,
      statusMessage: 'عدم دسترسی به سرویس محبوب‌ترین جستجوها'
    })
  }
})

export default popularSearchHandler
