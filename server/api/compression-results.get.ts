import { proxy } from './_utils/fetchProxy'
import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  return await proxy(event, 'http://localhost:3003/api/admin/media/compression-results')
}) 