import { defineEventHandler } from 'h3'
import { fetchGo } from '../../_utils/fetchGo'

type BackendSetting = {
  key?: string
  Key?: string
  value?: unknown
  Value?: unknown
}

type SocialSettings = Record<string, unknown>

function mapSettings(items: BackendSetting[]): SocialSettings {
  const settings: SocialSettings = {}
  const prefixedKeys = new Set<string>()

  for (const item of items) {
    const rawKey = (item?.key ?? item?.Key ?? '') as string
    if (!rawKey) {
      continue
    }

    const cleanKey = rawKey.replace(/^social-media\./, '')
    const isPrefixed = rawKey.startsWith('social-media.')
    const value = item?.value ?? item?.Value ?? ''

    if (isPrefixed) {
      prefixedKeys.add(cleanKey)
      settings[cleanKey] = value
    } else if (!prefixedKeys.has(cleanKey) && settings[cleanKey] === undefined) {
      settings[cleanKey] = value
    }
  }

  const rawCustomLinks = settings.custom_links ?? settings.customLinks
  if (typeof rawCustomLinks === 'string') {
    try {
      settings.customLinks = JSON.parse(rawCustomLinks)
    } catch (parseError) {
      console.warn('⚠️ Failed to parse custom_links JSON:', parseError)
      settings.customLinks = []
    }
  } else if (Array.isArray(rawCustomLinks)) {
    settings.customLinks = rawCustomLinks
  } else {
    settings.customLinks = []
  }

  return settings
}

export default defineEventHandler(async (event) => {
  try {
    const response = await fetchGo(event, '/api/settings', {
      method: 'GET'
    })

    const rawSettings = Array.isArray(response)
      ? response
      : Array.isArray((response as { data?: BackendSetting[] })?.data)
        ? (response as { data: BackendSetting[] }).data
        : []

    const settings = mapSettings(rawSettings)

    return {
      success: true,
      data: settings,
      message: 'تنظیمات شبکه‌های اجتماعی عمومی با موفقیت دریافت شد'
    }
  } catch (error: unknown) {
    console.error('❌ خطا در دریافت تنظیمات عمومی شبکه‌های اجتماعی:', error)
    return {
      success: false,
      data: {},
      message: (error as { data?: { message?: string }; message?: string }).data?.message || (error as { data?: { message?: string }; message?: string }).message || 'خطا در دریافت تنظیمات شبکه‌های اجتماعی'
    }
  }
})
