import { defineEventHandler, readBody } from 'h3'
import { proxy } from '../../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

// Retry یک job خطادار یا ایجاد مجدد job برای media_id
export default defineEventHandler(async (event) => {
     const body = await readBody(event)
     const config = useRuntimeConfig()
     return await proxy(event, `${config.public.goApiBase}/api/media/admin/image-seo/retry`, {
          method: 'POST',
          body: JSON.stringify(body)
     })
})


