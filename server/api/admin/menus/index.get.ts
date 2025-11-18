import { defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

interface Menu {
  id: number
  name: string
  slug: string
  [key: string]: unknown
}

interface MenuResponse {
  data?: Menu[]
  [key: string]: unknown
}

export default defineEventHandler(async (event) => {
  try {
    const response = await fetchGo(event, '/api/menus') as MenuResponse
    const menus = response?.data || []

    return {
      success: true,
      data: menus,
      message: 'منوها با موفقیت دریافت شدند'
    }
  } catch (error: unknown) {
    console.error('Error fetching menus:', error)

    return {
      success: false,
      data: [],
      message: 'خطا در دریافت منوها'
    }
  }
})