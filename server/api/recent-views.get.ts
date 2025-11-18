import { defineEventHandler, createError } from 'h3'

export default defineEventHandler(async (event) => {
  const { fetchGo } = await import('./_utils/fetchGo')

  const userId = event.context?.user?.id
  if (!userId) {
    throw createError({
      statusCode: 401,
      message: '˜ÇÑÈÑ ÇÍÑÇÒ åæíÊ äÔÏå ÇÓÊ'
    })
  }

  try {
    // ÇÑÓÇá ÏÑÎæÇÓÊ Èå ÓÑæíÓ Go ÈÑÇí ÏÑíÇİÊ áíÓÊ ÈÇÒÏíÏåÇí ÇÎíÑ ˜ÇÑÈÑ
    const response = await fetchGo(event, '/api/recent-views', { method: 'GET' })

    const data = Array.isArray((response as any)?.data)
      ? (response as any).data
      : Array.isArray(response) ? response : []

    return {
      success: true,
      data,
      count: data.length
    }
  } catch (error: any) {
    // ËÈÊ ÎØÇ ÈÑÇí ÈÑÑÓí æ ÇÑÓÇá íÇã ãäÇÓÈ Èå ˜ÇÑÈÑ
    console.error('ÎØÇ ÏÑ ÏÑíÇİÊ ÈÇÒÏíÏåÇí ÇÎíÑ:', error)
    throw createError({
      statusCode: error?.statusCode || 500,
      message: error?.statusMessage || 'ÎØÇ ÏÑ ÏÑíÇİÊ ÈÇÒÏíÏåÇí ÇÎíÑ'
    })
  }
})

