import type { H3Event } from 'h3'
import OpenAI from 'openai'

import { fetchGo } from '../../_utils/fetchGo'

declare const defineEventHandler: <T = unknown>(handler: (event: H3Event) => T | Promise<T>) => T
declare const readBody: (event: H3Event) => Promise<unknown>
// declare const createError: (options: { statusCode: number; message: string }) => Error

interface GenerateContentBody {
  prompt?: string
  model?: string
  wordCount?: number
  temperature?: number
  writingStyle?: string
  messages?: Array<{ role: string; content: string }>
}

interface ApiSettings {
  openai?: {
    api_key?: string
    api_url?: string
    enabled?: boolean
    default_model?: string
    temperature?: number
  }
}

export default defineEventHandler(async (event) => {
  try {
    const body = await readBody(event) as GenerateContentBody
    const { prompt, model, wordCount, temperature, writingStyle, messages } = body

    // console.log('🔍 درخواست تولید محتوا دریافت شد:', { prompt, model, wordCount, temperature, writingStyle })

    if (!prompt) {
      throw createError({
        statusCode: 400,
        message: 'پارامتر prompt الزامی است'
      })
    }

    // دریافت تنظیمات OpenAI
    // console.log('🔍 دریافت تنظیمات OpenAI...')
    const apiSettings = await fetchGo(event, '/api/admin/api-settings') as ApiSettings | null
    // console.log('📄 تنظیمات دریافت شده:', apiSettings ? 'موجود' : 'ناموجود')

    if (!apiSettings || !apiSettings.openai || !apiSettings.openai.api_key) {
      // console.error('❌ تنظیمات OpenAI یافت نشد')
      throw createError({
        statusCode: 400,
        message: 'تنظیمات OpenAI یافت نشد. لطفاً ابتدا تنظیمات API را در بخش تنظیمات ذخیره کنید.'
      })
    }

    const openAIConfig = apiSettings.openai
    const apiKey = openAIConfig.api_key
    const apiUrl = openAIConfig.api_url || 'https://api.openai.com/v1'

    // console.log('🔑 API Key موجود:', apiKey ? 'بله' : 'خیر')
    // console.log('🌐 API URL:', apiUrl)

    // بررسی فعال بودن OpenAI
    if (!openAIConfig.enabled) {
      // console.error('❌ OpenAI فعال نیست')
      throw createError({
        statusCode: 400,
        message: 'OpenAI فعال نیست. لطفاً ابتدا OpenAI را در تنظیمات فعال کنید.'
      })
    }

    // ایجاد OpenAI client
    const openai = new OpenAI({
      apiKey: apiKey,
      baseURL: apiUrl
    })

    // ساخت prompt کامل
    const systemPrompt = `تو یک نویسنده حرفه‌ای فارسی هستی. وظیفه تو تولید محتوای با کیفیت برای وبلاگ است.

تنظیمات:
- تعداد کلمات: ${wordCount || 500} کلمه
- سبک نوشتار: ${getWritingStyleText(writingStyle || '')}
- زبان: فارسی
- فرمت خروجی: JSON با فیلدهای زیر

فیلدهای مورد نیاز:


لطفاً محتوای تولید شده را در فرمت JSON برگردان.`

    // ساخت messages برای OpenAI
    type OpenAIMessageRole = 'system' | 'user' | 'assistant'

    interface Message {
      role: OpenAIMessageRole
      content: string
    }

    const safeMessages: Message[] = Array.isArray(messages) ? messages
      .filter((msg): msg is Message =>
        typeof msg === 'object' &&
        msg !== null &&
        'role' in msg &&
        'content' in msg &&
        (msg.role === 'system' || msg.role === 'user' || msg.role === 'assistant')
      )
      .map(msg => ({
        role: msg.role as OpenAIMessageRole,
        content: String(msg.content)
      }))
      : []

    const openAIMessages: OpenAI.Chat.Completions.ChatCompletionMessageParam[] = [
      { role: 'system', content: systemPrompt },
      ...safeMessages.map((msg: Message): OpenAI.Chat.Completions.ChatCompletionMessageParam => ({
        role: msg.role,
        content: msg.content
      })),
      { role: 'user', content: prompt }
    ]

    // console.log('📤 ارسال درخواست به OpenAI...')
    // console.log('  - مدل:', model || openAIConfig.default_model || 'gpt-3.5-turbo')
    // console.log('  - تعداد پیام‌ها:', openAIMessages.length)

    // ارسال درخواست به OpenAI با SDK
    const response = await openai.chat.completions.create({
      model: model || openAIConfig.default_model || 'gpt-3.5-turbo',
      messages: openAIMessages,
      max_tokens: Math.min((wordCount || 500) * 2, 4000), // حداکثر 4000 token
      temperature: temperature || openAIConfig.temperature || 0.7,
      top_p: 1,
      frequency_penalty: 0,
      presence_penalty: 0
    })

    // console.log('✅ پاسخ OpenAI دریافت شد')

    if (!response.choices || !response.choices[0] || !response.choices[0].message) {
      // console.error('❌ پاسخ نامعتبر از OpenAI:', response)
      throw createError({
        statusCode: 500,
        message: 'پاسخ نامعتبر از OpenAI'
      })
    }

    const aiResponse = response.choices[0].message.content || ''
    // console.log('📝 محتوای تولید شده:', aiResponse ? aiResponse.substring(0, 100) + '...' : '(خالی)')

    // تلاش برای parse کردن JSON
    let generatedContent
    try {
      // حذف کدهای markdown اگر وجود دارد
      const cleanResponse = aiResponse.replace(/```json\n?|\n?```/g, '').trim()
      generatedContent = JSON.parse(cleanResponse)
      // console.log('✅ JSON با موفقیت parse شد')
    } catch (parseError) {
      // console.warn('⚠️ خطا در parse کردن JSON، استفاده از محتوای متنی:', parseError)
      // اگر JSON نبود، محتوای متنی برگردان
      generatedContent = {
        title: 'مقاله تولید شده',
        excerpt: aiResponse ? aiResponse.substring(0, 200) + '...' : '',
        content: aiResponse,
        meta_title: 'مقاله تولید شده',
        meta_description: aiResponse ? aiResponse.substring(0, 160) + '...' : '',
        meta_keywords: 'مقاله, تولید شده',
        tags: ['مقاله'],
        slug: 'article-' + Date.now()
      }
    }

    // console.log('✅ محتوا با موفقیت تولید شد')
    return {
      content: aiResponse,
      generatedContent
    }

  } catch (error: unknown) {
    // console.error('❌ خطا در تولید محتوا:', error)

    // بررسی خطاهای خاص
    const errorWithStatus = error as { statusCode?: number; statusMessage?: string; message?: string }
    if (errorWithStatus?.statusCode === 401) {
      throw createError({
        statusCode: 401,
        message: 'API Key نامعتبر است. لطفاً کلید API صحیح را وارد کنید.'
      })
    } else if (errorWithStatus?.statusCode === 403) {
      throw createError({
        statusCode: 403,
        message: 'دسترسی به OpenAI محدود شده است. لطفاً تنظیمات API را بررسی کنید.'
      })
    } else if (errorWithStatus?.statusCode === 429) {
      throw createError({
        statusCode: 429,
        message: 'محدودیت تعداد درخواست. لطفاً کمی صبر کنید و دوباره تلاش کنید.'
      })
    } else if (errorWithStatus?.statusCode) {
      throw error
    }

    throw createError({
      statusCode: 500,
      message: 'خطا در تولید محتوا: ' + (errorWithStatus?.message || 'خطای نامشخص')
    })
  }
})

// تابع تبدیل سبک نوشتار به متن فارسی
function getWritingStyleText(style: string): string {
  const styles = {
    professional: 'حرفه‌ای و رسمی',
    casual: 'صمیمی و غیررسمی',
    academic: 'علمی و دانشگاهی',
    creative: 'خلاقانه و جذاب',
    technical: 'تکنیکی و تخصصی'
  }
  return styles[style as keyof typeof styles] || 'حرفه‌ای'
} 