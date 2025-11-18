import { defineEventHandler, readBody, createError } from 'h3'
import { getDatabase } from '../_utils/database.js'

export default defineEventHandler(async (event) => {
  try {
    const { id } = event.context.params
    if (!id) {
      throw createError({ statusCode: 400, message: 'Gateway ID is required.' })
    }

    const body = await readBody(event)
    if (typeof body.is_active === 'undefined') {
      throw createError({ statusCode: 400, message: 'is_active field is missing in request body.' })
    }

    const db = await getDatabase()

    // Convert truthy/falsy to boolean
    const isActive = !!body.is_active

    const result = await db.query(
      `UPDATE sms_gateways
       SET is_active = $1, updated_at = NOW()
       WHERE id = $2
       RETURNING id, name, type, is_active, updated_at;`,
      [isActive, id]
    )

    if (!result || result.length === 0) {
      throw createError({ statusCode: 404, message: 'Gateway not found or could not be updated.' })
    }

    return {
      status: 'success',
      data: result[0]
    }
  } catch (error) {
    console.error('Error in PATCH /api/sms-gateways/[id]:', error)
    throw createError({
      statusCode: error.statusCode || 500,
      message: error.message || error.statusMessage || 'Failed to update sms gateway'
    })
  }
}) 