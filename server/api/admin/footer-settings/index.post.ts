import { defineEventHandler, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

const normalizeFooterPayload = (payload: unknown) => {
     if (!payload || typeof payload !== 'object') {
          return payload
     }

     const normalized = { ...(payload as Record<string, unknown>) }

     if (Array.isArray(normalized.layers)) {
          normalized.layers = normalized.layers.map((layer: unknown) => {
               if (!layer || typeof layer !== 'object') {
                    return layer
               }

               const normalizedLayer = { ...(layer as Record<string, unknown>) }

               if (normalizedLayer.items !== undefined && typeof normalizedLayer.items !== 'string') {
                    try {
                         normalizedLayer.items = JSON.stringify(normalizedLayer.items)
                    } catch {
                         // console.error('Failed to stringify footer layer items:', error)
                         normalizedLayer.items = '[]'
                    }
               }

               if (normalizedLayer.boxWidths !== undefined && typeof normalizedLayer.boxWidths !== 'string') {
                    try {
                         normalizedLayer.boxWidths = JSON.stringify(normalizedLayer.boxWidths)
                    } catch {
                         // console.error('Failed to stringify footer layer boxWidths:', error)
                         delete normalizedLayer.boxWidths
                    }
               }

               if (normalizedLayer.box_widths !== undefined && typeof normalizedLayer.box_widths !== 'string') {
                    try {
                         normalizedLayer.box_widths = JSON.stringify(normalizedLayer.box_widths)
                    } catch {
                         // console.error('Failed to stringify footer layer box_widths:', error)
                         delete normalizedLayer.box_widths
                    }
               }

               return normalizedLayer
          })
     }

     return normalized
}

interface FooterCreateResponse {
     success: boolean
     data?: unknown
     message?: string
}

export default defineEventHandler(async (event): Promise<FooterCreateResponse> => {
     try {
          const body = await readBody(event)
          const payload = normalizeFooterPayload(body)

          // console.log('Creating admin footer...')

          let result: unknown
          try {
               result = await fetchGo(event, '/api/admin/footer-settings', {
                    method: 'POST',
                    body: payload
               })
          } catch (fetchErr: unknown) {
               const err = fetchErr as { data?: { message?: string; error?: string }; message?: string }
               // console.error('Admin footer creation failed:', {
               //      statusCode: err?.statusCode,
               //      status: err?.status,
               //      message: err?.message,
               //      data: err?.data
               // })
               return {
                    success: false,
                    message: err?.data?.message || err?.data?.error || err?.message || 'خطا در ایجاد فوتر'
               }
          }

          // console.log('Admin footer created successfully.')

          return {
               success: true,
               data: (result as { data?: unknown })?.data ?? result
          }
     } catch (error: unknown) {
          const err = error as { data?: { message?: string; error?: string }; message?: string }
          // console.error('Unexpected error during admin footer creation:', error)
          return {
               success: false,
               message: err?.data?.message || err?.data?.error || err?.message || 'خطا در ایجاد فوتر'
          }
     }
})
