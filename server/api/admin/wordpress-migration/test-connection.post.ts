import type { H3Event } from 'h3'
import { createError, defineEventHandler, readBody } from 'h3'

interface TestConnectionBody {
  siteUrl?: string
  consumerKey?: string
  consumerSecret?: string
}

interface WordPressResponse {
  name?: string
  description?: string
  url?: string
  namespaces?: string[]
}

interface WooCommerceResponse {
  environment?: {
    wp_version?: string
    version?: string
  }
  database?: {
    wc_database_version?: string
  }
}

export default defineEventHandler(async (event: H3Event) => {
  try {
    const body = await readBody(event) as TestConnectionBody
    const { siteUrl, consumerKey, consumerSecret } = body

    if (!siteUrl) {
      throw createError({
        statusCode: 400,
        message: 'آدرس سایت الزامی است'
      })
    }

    // تست اتصال به WooCommerce API
    if (consumerKey && consumerSecret) {
      try {
        const testUrl = `${siteUrl}/wp-json/wc/v3/system_status`
        const auth = Buffer.from(`${consumerKey}:${consumerSecret}`).toString('base64')

        const response = await fetch(testUrl, {
          headers: {
            'Authorization': `Basic ${auth}`,
            'Content-Type': 'application/json'
          }
        })

        if (!response.ok) {
          throw new Error(`HTTP ${response.status}: ${response.statusText}`)
        }

        const data = await response.json() as WooCommerceResponse

        return {
          success: true,
          message: 'اتصال به WooCommerce با موفقیت برقرار شد',
          data: {
            wordpress_version: data.environment?.wp_version || 'نامشخص',
            woocommerce_version: data.environment?.version || 'نامشخص',
            database: data.database?.wc_database_version || 'نامشخص'
          }
        }
      } catch (error: unknown) {
        const errorWithMessage = error as { message?: string }
        throw createError({
          statusCode: 500,
          message: `خطا در اتصال به WooCommerce: ${errorWithMessage.message || 'خطای نامشخص'}`
        })
      }
    }

    // تست اتصال ساده به سایت
    try {
      const response = await fetch(`${siteUrl}/wp-json`, {
        method: 'GET'
      })

      if (!response.ok) {
        throw new Error(`HTTP ${response.status}`)
      }

      const data = await response.json() as WordPressResponse

      return {
        success: true,
        message: 'اتصال به وردپرس با موفقیت برقرار شد',
        data: {
          name: data.name || 'نامشخص',
          description: data.description || '',
          url: data.url || siteUrl,
          namespaces: data.namespaces || []
        }
      }
    } catch (error: unknown) {
      const errorWithMessage = error as { message?: string }
      throw createError({
        statusCode: 500,
        message: `خطا در اتصال به وردپرس: ${errorWithMessage.message || 'خطای نامشخص'}`
      })
    }
  } catch (error: unknown) {
    console.error('خطا در تست اتصال:', error)

    const errorWithStatus = error as { statusCode?: number; message?: string }
    if (errorWithStatus.statusCode) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: errorWithStatus.message || 'خطا در تست اتصال'
    })
  }
})
