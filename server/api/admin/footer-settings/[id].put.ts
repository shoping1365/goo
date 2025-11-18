import { defineEventHandler, getRouterParam, readBody } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

const normalizeFooterPayload = (payload: any) => {
     if (!payload || typeof payload !== 'object') {
          return payload
     }

     const normalized: any = { ...payload }

     if (Array.isArray(normalized.layers)) {
          normalized.layers = normalized.layers.map((layer: any) => {
               if (!layer || typeof layer !== 'object') {
                    return layer
               }

               const normalizedLayer: any = { ...layer }

               if (normalizedLayer.items !== undefined && typeof normalizedLayer.items !== 'string') {
                    try {
                         normalizedLayer.items = JSON.stringify(normalizedLayer.items)
                    } catch (error) {
                         console.error('Failed to stringify footer layer items:', error)
                         normalizedLayer.items = '[]'
                    }
               }

               if (normalizedLayer.boxWidths !== undefined && typeof normalizedLayer.boxWidths !== 'string') {
                    try {
                         normalizedLayer.boxWidths = JSON.stringify(normalizedLayer.boxWidths)
                    } catch (error) {
                         console.error('Failed to stringify footer layer boxWidths:', error)
                         delete normalizedLayer.boxWidths
                    }
               }

               if (normalizedLayer.box_widths !== undefined && typeof normalizedLayer.box_widths !== 'string') {
                    try {
                         normalizedLayer.box_widths = JSON.stringify(normalizedLayer.box_widths)
                    } catch (error) {
                         console.error('Failed to stringify footer layer box_widths:', error)
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
  data?: any
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
    console.log('Updating admin footer:', id)

          let result: any
          try {
               result = await fetchGo(event, `/api/admin/footer-settings/${id}`, {
                    method: 'PUT',
                    body: payload
               })
          } catch (fetchErr: any) {
               console.error('Admin footer update failed:', {
                    statusCode: fetchErr?.statusCode,
                    status: fetchErr?.status,
                    message: fetchErr?.message,
                    data: fetchErr?.data
               })
               return {
                    success: false,
                    message: fetchErr?.data?.message || fetchErr?.data?.error || fetchErr?.message || 'خطا در به‌روزرسانی فوتر'
               }
          }

          console.log('Admin footer updated successfully.')

          return {
               success: true,
               data: result?.data ?? result
          }
  } catch (error: any) {
    console.error('Unexpected error during admin footer update:', error)
    return {
      success: false,
               message: error?.data?.message || error?.data?.error || error?.message || 'خطا در به‌روزرسانی فوتر'
    }
  }
})
