import { H3Event, createError, getRouterParam } from 'h3'
import { proxy } from '../../../../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

type MessageMethod = 'GET' | 'POST'

interface SessionMessageOptions {
  method: MessageMethod
  errorMessage: string
  logLabel: string
}

export async function handleSessionMessages(event: H3Event, options: SessionMessageOptions) {
  const sessionId = getRouterParam(event, 'id')
  if (!sessionId) {
    throw createError({ statusCode: 400, message: 'شناسه جلسه الزامی است' })
  }

  const config = useRuntimeConfig()
  const url = `${config.public.goApiBase}/api/chat/admin/sessions/${sessionId}/messages`

  try {
    return await proxy(event, url, { method: options.method })
  } catch (error) {
    console.error(options.logLabel, error)
    throw createError({ statusCode: 500, message: options.errorMessage })
  }
}
