import { defineEventHandler, createError, setResponseHeaders } from 'h3'

interface QAItem {
  id: number
  question: string
  answer: string
  status: string
  created_at: string
  updated_at: string
}

function convertToCSV(data: QAItem[]): string {
  const headers = ['شناسه', 'سوال', 'پاسخ', 'وضعیت', 'تاریخ ایجاد', 'تاریخ بروزرسانی']
  const rows = data.map(item => [
    item.id,
    `"${item.question.replace(/"/g, '""')}"`,
    `"${item.answer.replace(/"/g, '""')}"`,
    item.status,
    new Date(item.created_at).toLocaleDateString('fa-IR'),
    new Date(item.updated_at).toLocaleDateString('fa-IR')
  ])

  return [
    headers.join(','),
    ...rows.map(row => row.join(','))
  ].join('\n')
}

export default defineEventHandler(async (event) => {
  try {
    const config = useRuntimeConfig()
    const base = config.public.goApiBase
    // Get QA data from API
    const response = await fetch(`${base}/api/qa`)
    if (!response.ok) {
      throw new Error('Failed to fetch QA data')
    }
    const qaData = await response.json()

    // Convert to CSV
    const csv = convertToCSV(qaData)

    // Add BOM for Excel to recognize UTF-8
    const buffer = Buffer.from('\uFEFF' + csv, 'utf-8')

    // Set response headers
    setResponseHeaders(event, {
      'Content-Type': 'text/csv; charset=utf-8',
      'Content-Disposition': 'attachment; filename=qa-export.csv'
    })

    return buffer
  } catch (error) {
    console.error('Error exporting QA data:', error)
    throw createError({
      statusCode: 500,
      message: 'خطا در خروجی گرفتن داده‌ها'
    })
  }
})