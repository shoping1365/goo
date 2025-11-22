import { createError, defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const { getDatabase } = await import('../_utils/database')
    const db = await getDatabase()

    // Fetch user profile from database
    const userResult = await db.query(
      'SELECT id, name, email, mobile, profile_data FROM users WHERE id = $1',
      [event.context?.user?.id || 0]
    )
    
    if (!userResult || userResult.length === 0) {
      throw createError({
        statusCode: 404,
        message: 'User not found'
      })
    }
    
    const user = userResult[0] as { id: number; name: string | null; email: string | null; mobile: string | null; profile_data: unknown }
    
    return {
      id: user.id,
      name: user.name,
      email: user.email,
      mobile: user.mobile,
      profile_data: user.profile_data
    }

  } catch (error: unknown) {
    console.error('Error fetching user profile:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'Error fetching user profile'
    })
  }
})

