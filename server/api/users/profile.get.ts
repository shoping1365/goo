import { defineEventHandler, createError } from 'h3'

export default defineEventHandler(async (event) => {
  try {
    const { getDatabase } = await import('../_utils/database')
    const db = await getDatabase()

    // 
    const userResult = await db.query(
      'SELECT id, name, email, mobile, profile_data FROM users WHERE id = $1',
      [event.context?.user?.id || 0]
    )
    
    if (!userResult || userResult.length === 0) {
      throw createError({
        statusCode: 404,
        message: 'ò«—»— Ì«›  ‰‘œ'
      })
    }
    
    const userData = userResult[0] as any
    
    return {
      id: userData.id,
      name: userData.name,
      email: userData.email,
      mobile: userData.mobile,
      profile_data: userData.profile_data
    }

  } catch (error: any) {
    console.error('Œÿ« œ— œ—Ì«›  Å—Ê›«Ì· ò«—»—:', error)
    
    if (error.statusCode) {
      throw error
    }
    
    throw createError({
      statusCode: 500,
      message: 'Œÿ« œ— œ—Ì«›  Å—Ê›«Ì· ò«—»—'
    })
  }
})

