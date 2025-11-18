import { createError, defineEventHandler, readBody } from 'h3'
import { proxy } from '../_utils/fetchProxy'

declare const useRuntimeConfig: () => { public: { goApiBase: string } }

interface MenuAssignBody {
  menu_id?: number
  location_id?: number
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  const body = await readBody(event) as MenuAssignBody
  const config = useRuntimeConfig()

  if (!body.menu_id || !body.location_id) {
    throw createError({
      statusCode: 400,
      message: 'شناسه منو و جایگاه الزامی است'
    })
  }

  return await proxy(event, `${config.public.goApiBase}/api/menu-assign/menu/${body.menu_id}/location/${body.location_id}`, {
    method: 'DELETE'
  })
})