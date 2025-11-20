import { defineEventHandler, getRouterParam, readBody } from 'h3'
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

interface FooterUpdateResponse {
  success: boolean
  data?: unknown
  message?: string
}

export default defineEventHandler(async (event): Promise<FooterUpdateResponse> => {
  try {
    const id = getRouterParam(event, 'id')

    if (!id) {
      return {
        success: false,
        message: 'شناسه فوتر الزامی است'
      }
    }

    const body = await readBody(event)
    const payload = normalizeFooterPayload(body)
    // console.log('Updating admin footer:', id)

          let result: unknown
          try {
               result = await fetchGo(event, `/api/admin/footer-settings/${id}`, {
                    method: 'PUT',
                    body: payload
               })
          } catch (fetchErr: unknown) {
               const err = fetchErr as { data?: { message?: string; error?: string }; message?: string }
               // console.error('Admin footer update failed:', {
               //      statusCode: err?.statusCode,
               //      status: err?.status,
               //      message: err?.message,
               //      data: err?.data
               // })
               return {
                    success: false,
                    message: err?.data?.message || err?.data?.error || err?.message || 'خطا در به‌روزرسانی فوتر'
               }
          }

          // console.log('Admin footer updated successfully.')

          return {
               success: true,
               data: (result as { data?: unknown })?.data ?? result
          }
  } catch (error: unknown) {
    // console.error('Unexpected error during admin footer update:', error)
    return {
      success: false,
               message: error?.data?.message || error?.data?.error || error?.message || 'خطا در به‌روزرسانی فوتر'
    }
  }
})
