import type { H3Event } from 'h3'
import { defineEventHandler } from 'h3'
import { proxy } from '../../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event: H3Event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase

  // fetchProxy automatically forwards cookies and creates Authorization header
  return proxy(event, `${base}/api/chat/admin/ai-bots`)
})




