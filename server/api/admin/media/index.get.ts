import { defineEventHandler } from 'h3'

interface MediaItem {
  id: string
  url: string
  alt: string
  title: string
  [key: string]: unknown
}

interface MediaListResponse {
  data: MediaItem[]
  [key: string]: unknown
}

declare const $fetch: <T = unknown>(url: string) => Promise<T>

export default defineEventHandler(async (): Promise<{ data: MediaItem[] }> => {
  const response = await $fetch<MediaListResponse>('/api/media/list')
  return { data: response.data || [] }
})