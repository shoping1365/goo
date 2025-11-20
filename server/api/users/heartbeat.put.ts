import { createError, defineEventHandler, getCookie } from 'h3'
import { goApiFetch } from '~/server/utils/goApi'

export default defineEventHandler(async (event) => {
  try {
    const accessToken = getCookie(event, 'access_token')
    const refreshToken = getCookie(event, 'refresh_token')

    if (!accessToken && !refreshToken) {
      return { success: true, message: 'User not logged in, heartbeat ignored' }
    }

    try {
      const data = await goApiFetch(event, '/api/users/heartbeat', {
        method: 'PUT',
        body: { timestamp: Date.now() }
      })

      return data
    } catch {
      return { success: true, message: 'Heartbeat sent to Go Backend' }
    }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } catch (error: any) {
    // console.error('خطا در heartbeat:', error)

    if (error && typeof error === 'object' && 'statusCode' in error) {
      throw error
    }

    throw createError({ statusCode: 500, message: 'خطا در heartbeat' })
  }
})