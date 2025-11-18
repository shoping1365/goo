import { defineEventHandler } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  return await proxy(event, `${config.public.goApiBase}/api/admin/gift-cards/dashboard/stats`)
}) 