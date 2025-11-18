<template>
  <footer v-if="shouldShowFooter && activeFooter && activeFooter.layers" :style="footerStyle as any" class="footer">
    <div v-for="layer in activeFooter.layers" :key="layer.id" class="footer-layer" :style="getLayerStyle(layer)">
      <template
        v-for="(item, idx) in getLayerItems(layer)"
        :key="item.id || idx"
      >
        <span :style="getItemStyle(item)" class="footer-item-wrapper">
          <component
            v-if="resolveWidgetComponent(item)"
            :is="resolveWidgetComponent(item)"
            v-bind="getWidgetProps(item)"
          />

          <NuxtLink
            v-else-if="item.path"
            :to="item.path"
            class="footer-item-link"
          >
            {{ item.text || item.name || item.type || item }}
          </NuxtLink>

          <div v-else-if="typeof item === 'string'" class="footer-item">
            {{ item }}
          </div>
        </span>

        <!-- جداکننده بین آیتم‌ها -->
        <div
          v-if="layer.showSeparator && idx < getLayerItems(layer).length - 1"
          class="footer-separator"
          :style="getSeparatorStyle(layer)"
        ></div>
      </template>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { useRoute, useState } from 'nuxt/app'
import { computed, onMounted, ref, watch } from 'vue'
import { usePublicFooters } from '~/composables/usePublicFooters'

const route = useRoute()
const { loadFooters, getActiveFooter, getFooterForPage } = usePublicFooters()
const activeFooter = ref<any>(null)
const footerRef = ref<HTMLElement>()
const socialMediaSettings = ref<any>(null)
const companySettings = ref<any>(null)

// بارگذاری تنظیمات شبکه‌های اجتماعی
const getSocialMediaSettings = async () => {
  if (socialMediaSettings.value) {
    return socialMediaSettings.value
  }

  const fallbackResult = {
    socials: [] as any[],
    links: [] as any[],
    customLinks: [] as any[]
  }

  let lastFetchError: unknown = null

  try {
  const endpoints = [
    '/api/public/settings/social-media',
    '/nuxt/admin/settings/social-media'
  ]
    let response: any = null

    for (const endpoint of endpoints) {
      try {
        const result = await $fetch<any>(endpoint)
        if (result && result.success && result.data) {
          response = result
          break
        }
      } catch {
        // Ignore fetch errors and try next endpoint
      }
    }

  if (response && response.success && response.data) {
      const settings = response.data
      
      // ساخت آرایه socials برای FooterWidgetSocial
      const socials: any[] = []
      let parsedCustomLinks: any[] = []
      const platformLabels: Record<string, string> = {
        instagram: 'اینستاگرام',
        telegram: 'تلگرام',
        whatsapp: 'واتساپ',
        twitter: 'توییتر',
        linkedin: 'لینکدین',
        facebook: 'فیسبوک',
        youtube: 'یوتیوب',
        aparat: 'آپارات',
        rubika: 'روبیکا',
        eitaa: 'ایتا',
        bale: 'بله',
        tiktok: 'تیک‌تاک',
        pinterest: 'پینترست',
        discord: 'دیسکورد',
        github: 'گیت‌هاب'
      }
      
      Object.keys(platformLabels).forEach(platform => {
        const enabledRaw = settings[`${platform}_enabled`]
        const urlRaw = settings[platform]

        const url = typeof urlRaw === 'string' ? urlRaw.trim() : (urlRaw ?? '')
  const hasUrl = typeof url === 'string' && url.length > 0
        const enabledNormalized = typeof enabledRaw === 'string'
          ? enabledRaw.trim().toLowerCase()
          : enabledRaw
        const isEnabled = enabledNormalized === true
          || enabledNormalized === 'true'
          || enabledNormalized === '1'
          || enabledNormalized === 'yes'
          || enabledNormalized === 'on'
          || enabledNormalized === 'active'
        const hasExplicitEnabled = enabledRaw !== undefined && enabledRaw !== null && `${enabledRaw}` !== ''
        const shouldInclude = isEnabled || (!hasExplicitEnabled && hasUrl)

        if (shouldInclude) {
          socials.push({
            platform,
            label: platformLabels[platform],
            href: hasUrl ? url : '#',
            url: hasUrl ? url : '',
            enabled: Boolean(isEnabled || hasUrl)
          })
        }
      })

      const rawCustomLinks = settings.custom_links ?? settings.customLinks
      if (rawCustomLinks) {
        if (typeof rawCustomLinks === 'string') {
          try {
            parsedCustomLinks = JSON.parse(rawCustomLinks)
          } catch {
            parsedCustomLinks = []
          }
        } else if (Array.isArray(rawCustomLinks)) {
          parsedCustomLinks = rawCustomLinks
        }

        if (Array.isArray(parsedCustomLinks)) {
          parsedCustomLinks.forEach((link: any, index: number) => {
            const rawHref = link?.href ?? link?.url
            const href = typeof rawHref === 'string' ? rawHref.trim() : ''
            if (!href) {
              return
            }
            socials.push({
              platform: link?.platform || `custom${index + 1}`,
              label: link?.title || link?.label || `لینک ${index + 1}`,
              href,
              url: href,
              enabled: true
            })
          })
        }
      }
      
      const result = {
        socials,
        links: socials,
        customLinks: Array.isArray(parsedCustomLinks) ? parsedCustomLinks : []
      }
      socialMediaSettings.value = result
      return result
    }
  } catch {
    // ignore social settings load errors
  }
  socialMediaSettings.value = fallbackResult
  return fallbackResult
}

const PHONE_SPLIT_REGEX = /[\n,;|،]+/

const normalizeText = (value: string) => value.replace(/\s+/g, ' ').trim()

function safeParseJSON(value: string) {
  try {
    return JSON.parse(value)
  } catch {
    return null
  }
}

function normalizePhoneValue(phone: string): string {
  return phone.replace(/\s+/g, ' ').trim()
}

function parsePhoneList(value: unknown): string[] {
  if (value === null || value === undefined) {
    return []
  }

  if (Array.isArray(value)) {
    return value.flatMap(item => parsePhoneList(item))
  }

  if (typeof value === 'number') {
    return [normalizePhoneValue(String(value))]
  }

  if (typeof value === 'string') {
    const trimmed = value.trim()
    if (!trimmed) return []

    if ((trimmed.startsWith('[') && trimmed.endsWith(']')) || (trimmed.startsWith('{') && trimmed.endsWith('}'))) {
      const parsed = safeParseJSON(trimmed)
      if (parsed !== null) {
        return parsePhoneList(parsed)
      }
    }

    return trimmed
      .split(PHONE_SPLIT_REGEX)
      .map(part => normalizePhoneValue(part))
      .filter(Boolean)
  }

  if (typeof value === 'object') {
    return Object.values(value as Record<string, unknown>).flatMap(val => parsePhoneList(val))
  }

  return []
}

function normalizeArrayInput(raw: unknown): any[] {
  if (Array.isArray(raw)) {
    return raw
  }

  if (typeof raw === 'string') {
    const trimmed = raw.trim()
    if (!trimmed) return []
    try {
      const parsed = JSON.parse(trimmed)
      if (Array.isArray(parsed)) {
        return parsed
      }
      if (parsed && typeof parsed === 'object') {
        return Object.values(parsed as Record<string, any>)
      }
    } catch {
      return []
    }
  }

  if (raw && typeof raw === 'object') {
    return Object.values(raw as Record<string, any>)
  }

  return []
}

function normalizeSocialEntries(raw: unknown): any[] {
  const list = normalizeArrayInput(raw)
  const normalized: any[] = []

  list.forEach((entry, index) => {
    if (!entry && entry !== 0) {
      return
    }

    if (typeof entry === 'string') {
      const href = entry.trim()
      if (!href || href === '#') {
        return
      }
      normalized.push({
        id: `social-${index}`,
        platform: '',
        label: href,
        title: href,
        href,
        url: href,
        enabled: true,
        visible: true,
        openInNewTab: true
      })
      return
    }

    if (typeof entry !== 'object') {
      return
    }

    const record = entry as Record<string, any>
    const href = typeof record.href === 'string'
      ? record.href.trim()
      : typeof record.url === 'string'
        ? record.url.trim()
        : ''

    if (!href || href === '#') {
      return
    }

    if (record.enabled === false || record.visible === false) {
      return
    }

    const platformRaw = record.platform ?? record.type ?? record.id ?? ''
    const platform = typeof platformRaw === 'string' ? platformRaw.trim() : ''
    const labelRaw = record.label ?? record.title ?? (platform || href)
    const label = typeof labelRaw === 'string' ? labelRaw.trim() : platform || href

    normalized.push({
      ...record,
      id: record.id ?? `social-${platform || index}`,
      platform,
      label,
      title: record.title ?? label,
      href,
      url: href,
      enabled: true,
      visible: true,
      openInNewTab: record.openInNewTab !== false
    })
  })

  return normalized
}

function extractAddressFromLocation(location: Record<string, any>): string {
  const candidates: string[] = []

  const rawAddress = location?.address ?? location?.Address ?? location?.fullAddress ?? location?.full_address ?? location?.addressText ?? location?.address_text
  if (typeof rawAddress === 'string' && rawAddress.trim()) {
    candidates.push(rawAddress)
  } else if (rawAddress && typeof rawAddress === 'object') {
    const parts = [
      rawAddress.line1 ?? rawAddress.addressLine1 ?? rawAddress.street ?? rawAddress.street1,
      rawAddress.line2 ?? rawAddress.addressLine2 ?? rawAddress.street2,
      rawAddress.neighborhood ?? rawAddress.neighbourhood,
      rawAddress.city ?? rawAddress.City,
      rawAddress.province ?? rawAddress.Province ?? rawAddress.state ?? rawAddress.State,
      rawAddress.postalCode ?? rawAddress.PostalCode ?? rawAddress.zipCode ?? rawAddress.zip_code ?? rawAddress.zip
    ].filter(Boolean)

    if (parts.length) {
      candidates.push(parts.join('، '))
    }
  }

  const flatParts = [
    location.addressLine1 ?? location.address_line1 ?? location.line1,
    location.addressLine2 ?? location.address_line2 ?? location.line2,
    location.street ?? location.Street,
    location.neighborhood ?? location.neighbourhood,
    location.city ?? location.City,
    location.province ?? location.Province ?? location.state ?? location.State,
    location.postalCode ?? location.postal_code ?? location.zipCode ?? location.zip_code ?? location.zip
  ].filter(Boolean)

  if (flatParts.length) {
    candidates.push(flatParts.join('، '))
  }

  const addressCandidate = candidates.find(candidate => typeof candidate === 'string' && candidate.trim())
  return addressCandidate ? normalizeText(addressCandidate) : ''
}

function pickFirstEmail(value: unknown): string | null {
  if (!value) return null

  if (typeof value === 'string') {
    const trimmed = value.trim()
    return trimmed || null
  }

  if (Array.isArray(value)) {
    for (const item of value) {
      const found = pickFirstEmail(item)
      if (found) return found
    }
  } else if (typeof value === 'object') {
    for (const item of Object.values(value as Record<string, unknown>)) {
      const found = pickFirstEmail(item)
      if (found) return found
    }
  }

  return null
}

type NormalizedLocation = {
  id: string | number
  title: string
  address: string
  phones: string[]
  email?: string
}

function normalizeLocationsInput(raw: unknown): NormalizedLocation[] {
  if (!raw && raw !== 0) {
    return []
  }

  let list: any[] = []

  if (typeof raw === 'string') {
    const parsed = safeParseJSON(raw)
    if (Array.isArray(parsed)) {
      list = parsed
    } else if (parsed && typeof parsed === 'object') {
      list = Object.values(parsed)
    } else if (raw.trim()) {
      list = [raw.trim()]
    }
  } else if (Array.isArray(raw)) {
    list = raw
  } else if (typeof raw === 'object') {
    list = Object.values(raw as Record<string, unknown>)
  }

  const result: NormalizedLocation[] = []

  list.forEach((entry, index) => {
    if (entry === null || entry === undefined) return

    if (typeof entry === 'string') {
      const address = normalizeText(entry)
      if (address) {
        result.push({
          id: `loc-${index}`,
          title: '',
          address,
          phones: []
        })
      }
      return
    }

    if (typeof entry !== 'object') return

    const rawObject = entry as Record<string, any>
    const id = rawObject.id ?? rawObject.Id ?? rawObject.locationId ?? rawObject.location_id ?? rawObject.code ?? `loc-${index}`
    const titleRaw = rawObject.title ?? rawObject.Title ?? rawObject.name ?? rawObject.Name ?? ''
    const title = typeof titleRaw === 'string' ? normalizeText(titleRaw) : ''
    const address = extractAddressFromLocation(rawObject)

    const phoneCollector = new Set<string>()
    parsePhoneList(rawObject.phones ?? rawObject.phone ?? rawObject.phoneNumbers ?? rawObject.phone_numbers ?? rawObject.tel ?? rawObject.telephone ?? rawObject.landline).forEach(value => phoneCollector.add(value))
    parsePhoneList(rawObject.mobiles ?? rawObject.mobile ?? rawObject.mobilePhones ?? rawObject.mobile_numbers).forEach(value => phoneCollector.add(value))
    parsePhoneList(rawObject.whatsapp ?? rawObject.whatsappNumber ?? rawObject.contact?.whatsapp).forEach(value => phoneCollector.add(value))
    parsePhoneList(rawObject.contact?.phones ?? rawObject.contact?.phone).forEach(value => phoneCollector.add(value))
    parsePhoneList(rawObject.contact?.mobiles ?? rawObject.contact?.mobile).forEach(value => phoneCollector.add(value))

    const phoneList = Array.from(phoneCollector).map(normalizePhoneValue).filter(Boolean)
    const email = pickFirstEmail(rawObject.email ?? rawObject.contact?.email ?? rawObject.emails ?? rawObject.contact?.emails ?? rawObject.contact?.emailAddress)

    if (!title && !address && phoneList.length === 0 && !email) {
      return
    }

    const normalizedLocation: NormalizedLocation = {
      id,
      title,
      address,
      phones: phoneList
    }

    if (email) {
      normalizedLocation.email = email
    }

    result.push(normalizedLocation)
  })

  return result
}

function isLikelyMobileNumber(phone: string): boolean {
  const digits = phone.replace(/\D+/g, '')
  if (!digits) return false

  if (digits.startsWith('0098')) {
    return digits.slice(4).startsWith('9')
  }

  if (digits.startsWith('98')) {
    return digits.slice(2).startsWith('9')
  }

  return digits.startsWith('09')
}

// بارگذاری تنظیمات شرکت
const getCompanySettings = async () => {
  if (companySettings.value) {
    return companySettings.value
  }

  try {
    const endpoints = [
      { url: '/api/shop-settings' },
      { url: '/api/admin/shop-settings' },
      { url: '/api/settings', params: { keys: 'company_phone,footer_phone_number,header_phone_number,company_address,company_email,working_hours' } },
      { url: '/api/admin/settings', params: { keys: 'company_phone,footer_phone_number,header_phone_number,company_address,company_email,working_hours' } }
    ]
    
    for (const endpoint of endpoints) {
      try {
        const response = await $fetch<any>(endpoint.url, endpoint.params ? { params: endpoint.params } : {})
        
        if (response) {
          let settingsObj: any = {}
          
          // تبدیل آرایه به object
          if (response?.status === 'success' && response?.data && typeof response.data === 'object') {
            settingsObj = response.data
          } else if (response.success && Array.isArray(response.data)) {
            response.data.forEach((item: any) => {
              if (item.key && item.value !== undefined) {
                settingsObj[item.key] = item.value
              }
            })
          } else if (Array.isArray(response.data)) {
            response.data.forEach((item: any) => {
              if (item.key && item.value !== undefined) {
                settingsObj[item.key] = item.value
              }
            })
          } else if (response.data && typeof response.data === 'object') {
            settingsObj = response.data
          } else if (Array.isArray(response)) {
            response.forEach((item: any) => {
              if (item.key && item.value !== undefined) {
                settingsObj[item.key] = item.value
              }
            })
          }
          
          // استخراج تلفن‌ها - از تنظیمات shop
          const phones: string[] = []
          
          // از تنظیمات shop-settings
          if (settingsObj.phones) {
            if (Array.isArray(settingsObj.phones)) {
              phones.push(...settingsObj.phones)
            } else if (typeof settingsObj.phones === 'string') {
              try {
                const parsed = JSON.parse(settingsObj.phones)
                if (Array.isArray(parsed)) {
                  phones.push(...parsed)
                } else {
                  phones.push(settingsObj.phones)
                }
              } catch {
                phones.push(settingsObj.phones)
              }
            }
          }
          
          // Fallback به key های قدیمی
          if (phones.length === 0) {
            if (settingsObj.company_phone) phones.push(settingsObj.company_phone)
            if (settingsObj.footer_phone_number) phones.push(settingsObj.footer_phone_number)
            if (settingsObj.header_phone_number) phones.push(settingsObj.header_phone_number)
            if (settingsObj.phone) phones.push(settingsObj.phone)
            if (settingsObj.phone_number) phones.push(settingsObj.phone_number)
          }
          
          // استخراج موبایل‌های مدیر
          const mobiles: string[] = []
          if (settingsObj.adminPhones) {
            if (Array.isArray(settingsObj.adminPhones)) {
              mobiles.push(...settingsObj.adminPhones)
            } else if (typeof settingsObj.adminPhones === 'string') {
              try {
                const parsed = JSON.parse(settingsObj.adminPhones)
                if (Array.isArray(parsed)) {
                  mobiles.push(...parsed)
                } else {
                  mobiles.push(settingsObj.adminPhones)
                }
              } catch {
                mobiles.push(settingsObj.adminPhones)
              }
            }
          }

          if (mobiles.length === 0) {
            if (Array.isArray(settingsObj.mobiles)) {
              mobiles.push(...settingsObj.mobiles)
            } else if (typeof settingsObj.mobiles === 'string' && settingsObj.mobiles.trim()) {
              mobiles.push(settingsObj.mobiles)
            }
            if (settingsObj.mobile_phone) mobiles.push(settingsObj.mobile_phone)
            if (settingsObj.mobile) mobiles.push(settingsObj.mobile)
          }
          
          // استخراج آدرس - از تنظیمات shop
          let address = settingsObj.address || 
                        settingsObj.company_address || 
                        settingsObj.footer_address || 
                        settingsObj.site_address || ''
          
          let secondaryAddress = settingsObj.secondaryAddress ||
                                   settingsObj.address_secondary ||
                                   settingsObj.second_address ||
                                   settingsObj.address2 || ''
          
          // استخراج ایمیل - از تنظیمات shop
          let email = settingsObj.email || 
                      settingsObj.company_email || 
                      settingsObj.footer_email || 
                      settingsObj.site_email || 
                      settingsObj.contact_email || ''

          const aboutDescription = settingsObj.shortDescription ||
                                   settingsObj.short_description ||
                                   settingsObj.aboutDescription ||
                                   settingsObj.about_description ||
                                   settingsObj.about ||
                                   settingsObj.about_us ||
                                   settingsObj.footerAbout ||
                                   settingsObj.footer_about ||
                                   settingsObj.description ||
                                   ''
          const normalizedAbout = typeof aboutDescription === 'string' ? aboutDescription.trim() : ''
          
          // استخراج ساعات کاری
          let workingSchedules: string[] = []
          if (settingsObj.workingHours) {
            try {
              if (typeof settingsObj.workingHours === 'string') {
                try {
                  workingSchedules = JSON.parse(settingsObj.workingHours)
                } catch {
                  workingSchedules = [settingsObj.workingHours]
                }
              } else if (Array.isArray(settingsObj.workingHours)) {
                workingSchedules = settingsObj.workingHours
              } else {
                workingSchedules = [String(settingsObj.workingHours)]
              }
            } catch {
              workingSchedules = []
            }
          }
          
          // Fallback به working_hours
          if (workingSchedules.length === 0 && settingsObj.working_hours) {
            try {
              if (typeof settingsObj.working_hours === 'string') {
                try {
                  workingSchedules = JSON.parse(settingsObj.working_hours)
                } catch {
                  workingSchedules = [settingsObj.working_hours]
                }
              } else if (Array.isArray(settingsObj.working_hours)) {
                workingSchedules = settingsObj.working_hours
              } else {
                workingSchedules = [String(settingsObj.working_hours)]
              }
            } catch {
              workingSchedules = []
            }
          }
          
          const rawLocations = settingsObj.locations ??
                               settingsObj.addresses ??
                               settingsObj.contactLocations ??
                               settingsObj.branches ??
                               settingsObj.shop_locations
          const normalizedLocations = normalizeLocationsInput(rawLocations)

          if (!address && normalizedLocations.length > 0) {
            address = normalizedLocations[0].address || address
          }

          if (!secondaryAddress && normalizedLocations.length > 1) {
            secondaryAddress = normalizedLocations[1]?.address || secondaryAddress
          }

          if (!email) {
            const locationWithEmail = normalizedLocations.find(location => location.email)
            if (locationWithEmail?.email) {
              email = locationWithEmail.email
            }
          }

          const basePhones = parsePhoneList(phones)
          const baseMobiles = parsePhoneList(mobiles)
          const locationPhones = normalizedLocations.flatMap(location => location.phones || [])
          const combinedPhones = Array.from(new Set([...basePhones, ...locationPhones])).map(normalizePhoneValue).filter(Boolean)
          const combinedMobiles = Array.from(new Set([
            ...baseMobiles,
            ...locationPhones.filter(isLikelyMobileNumber)
          ])).map(normalizePhoneValue).filter(Boolean)

          const result = {
            phones: combinedPhones,
            mobiles: combinedMobiles,
            address: address,
            secondaryAddress,
            email: email,
            about: normalizedAbout,
            workingSchedules: Array.isArray(workingSchedules) ? workingSchedules.filter(s => s && s.trim()) : [],
            locations: normalizedLocations
          }

          companySettings.value = result
          return result
        }
      } catch {
        // try next endpoint when loading company settings
      }
    }
  } catch {
    // ignore company settings load errors
  }
  
  return { phones: [], mobiles: [], address: '', secondaryAddress: '', email: '', about: '', workingSchedules: [], locations: [] }
}

// بررسی اینکه آیا فوتر باید نمایش داده شود یا نه
const shouldShowFooter = computed(() => {
  // استفاده از state میدل‌ور
  const footerDisplayState = useState('footerDisplayState', () => ({
    shouldShow: true,
    footerData: null
  }))

  // اگر state میدل‌ور موجود باشد، از آن استفاده کن
  if (footerDisplayState.value && footerDisplayState.value.shouldShow !== undefined) {
    return footerDisplayState.value.shouldShow
  }

  // اگر state میدل‌ور موجود نباشد، از منطق قبلی استفاده کن
  if (!activeFooter.value) return false

  try {
    const currentPath = route.path
    const footerForPage = getFooterForPage(currentPath)
    
    // اگر فوتر برای این صفحه وجود داشته باشد، نمایش بده
    return !!footerForPage
  } catch {
    return false
  }
})

const footerStyle = computed(() => {
  if (!activeFooter.value || !activeFooter.value.layers?.length) {
    return {}
  }

  const totalHeight = activeFooter.value.layers.reduce((total, layer) => total + (layer.height || 0), 0)

  return totalHeight > 0
    ? { minHeight: `${totalHeight}px` }
    : {}
})

function getLayerStyle(layer: any) {
  // سبک پایه فوتر را تعریف می‌کنیم و سپس هر استایل سفارشی درج‌شده در لایه را روی آن اعمال می‌کنیم
  const baseStyle: Record<string, any> = {
    minHeight: (layer.height ?? 50) + 'px',
    width: layer.width ? layer.width + '%' : '100%',
    display: 'flex',
    alignItems: 'center',
    justifyContent: getLayerJustifyContent(layer),
    direction: layer.direction || 'rtl',
    textAlign: layer.direction === 'ltr' ? ('left' as const) : ('right' as const),
    opacity: layer.opacity || 1,
    transition: 'all 0.3s ease',
    color: '#ffffff',
    overflow: 'hidden'
  }

  // اضافه کردن سایه اگر فعال باشد
  if (layer.enableShadow) {
    const shadowIntensity = layer.shadowIntensity || 'md'
    const shadowDirection = layer.shadowDirection || 'top'
    
    // تعریف انواع سایه
    const shadows = {
      sm: {
        top: '0 -2px 4px -1px rgba(0, 0, 0, 0.06)',
        bottom: '0 2px 4px -1px rgba(0, 0, 0, 0.06)',
        both: '0 -2px 4px -1px rgba(0, 0, 0, 0.06), 0 2px 4px -1px rgba(0, 0, 0, 0.06)'
      },
      md: {
        top: '0 -4px 6px -1px rgba(0, 0, 0, 0.1), 0 -2px 4px -1px rgba(0, 0, 0, 0.06)',
        bottom: '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
        both: '0 -4px 6px -1px rgba(0, 0, 0, 0.1), 0 4px 6px -1px rgba(0, 0, 0, 0.1)'
      },
      lg: {
        top: '0 -10px 15px -3px rgba(0, 0, 0, 0.1), 0 -4px 6px -2px rgba(0, 0, 0, 0.05)',
        bottom: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
        both: '0 -10px 15px -3px rgba(0, 0, 0, 0.1), 0 10px 15px -3px rgba(0, 0, 0, 0.1)'
      },
      xl: {
        top: '0 -20px 25px -5px rgba(0, 0, 0, 0.1), 0 -10px 10px -5px rgba(0, 0, 0, 0.04)',
        bottom: '0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)',
        both: '0 -20px 25px -5px rgba(0, 0, 0, 0.1), 0 20px 25px -5px rgba(0, 0, 0, 0.1)'
      }
    }
    
    baseStyle.boxShadow = shadows[shadowIntensity]?.[shadowDirection] || shadows.md.top
    baseStyle.position = 'relative'
    baseStyle.zIndex = 10
  }

  // اضافه کردن border اگر فعال باشد
  if (layer.enableBorder) {
    const borderColor = layer.borderColor || '#e5e7eb'
    const borderWidth = `${layer.borderWidth || 1}px`
    const borderStyle = layer.borderStyle || 'solid'
    const borderPosition = layer.borderPosition || 'all'
    
    switch (borderPosition) {
      case 'all':
        baseStyle.border = `${borderWidth} ${borderStyle} ${borderColor}`
        break
      case 'top':
        baseStyle.borderTop = `${borderWidth} ${borderStyle} ${borderColor}`
        break
      case 'bottom':
        baseStyle.borderBottom = `${borderWidth} ${borderStyle} ${borderColor}`
        break
      case 'left':
        baseStyle.borderLeft = `${borderWidth} ${borderStyle} ${borderColor}`
        break
      case 'right':
        baseStyle.borderRight = `${borderWidth} ${borderStyle} ${borderColor}`
        break
      case 'top-bottom':
        baseStyle.borderTop = `${borderWidth} ${borderStyle} ${borderColor}`
        baseStyle.borderBottom = `${borderWidth} ${borderStyle} ${borderColor}`
        break
    }
  }

  // لایه ممکن است یک آبجکت استایل یا یک رشته JSON باشد
  let customStyle: Record<string, any> = {}
  if (layer.style) {
    if (typeof layer.style === 'string') {
      try {
        customStyle = JSON.parse(layer.style)
        // گاهی اوقات به صورت اشتباه به جای height از totalHeight استفاده شده
        if (!('height' in customStyle) && typeof customStyle.totalHeight !== 'undefined') {
          customStyle.height = customStyle.totalHeight
        }
      } catch {
        customStyle = {}
      }
    } else if (typeof layer.style === 'object') {
      customStyle = layer.style
    }
  }

  return { ...baseStyle, ...customStyle }
}

/**
 * تعیین چینش لایه بر اساس آیتم‌ها و تنظیمات
 */
function getLayerJustifyContent(layer: any): string {
  // اگر لایه دارای آیتم‌هایی با چینش مشخص باشد
  if (layer.items && Array.isArray(layer.items)) {
    try {
      const items = typeof layer.items === 'string' ? JSON.parse(layer.items) : layer.items

      // بررسی اینکه آیا همه آیتم‌ها چینش یکسانی دارند
      if (items.length > 0) {
        const firstItem = items[0]
        const firstAlign = firstItem?.align || 'center'

        // اگر همه آیتم‌ها چینش یکسانی دارند
        const allSameAlign = items.every((item: any) => (item?.align || 'center') === firstAlign)

        if (allSameAlign) {
          switch (firstAlign) {
            case 'left':
              return 'flex-start'
            case 'center':
              return 'center'
            case 'right':
              return 'flex-end'
            default:
              return 'center'
          }
        }

        // اگر چینش‌ها متفاوت باشند، از space-between استفاده می‌کنیم
        return 'space-between'
      }
    } catch {
      // ignore layer alignment errors
    }
  }

  // اگر آیتم‌ها وجود نداشته باشند، از direction استفاده می‌کنیم
  return layer.direction === 'ltr' ? 'flex-start' : 'flex-end'
}

function getLayerItems(layer) {
  if (!layer || !layer.items) return []
  try {
    const rawItems = typeof layer.items === 'string' ? JSON.parse(layer.items) : layer.items

    return Array.isArray(rawItems)
      ? rawItems.map((raw: any) => normalizeItem(raw))
      : []
  } catch {
    return []
  }
}

function normalizeItem(raw: any) {
  if (raw && typeof raw === 'object') {
    const resolvedText = (() => {
      if (typeof raw.text === 'string' && raw.text.trim().length) {
        return raw.text
      }
      if (raw.props && typeof raw.props.text === 'string' && raw.props.text.trim().length) {
        return raw.props.text
      }
      const fallback = raw.copyrightText || raw.title || raw.name || raw.content || raw.value
      return typeof fallback === 'string' ? fallback : ''
    })()

    const normalizedProps: Record<string, any> = {
      ...(raw.props || {})
    }

    if (resolvedText && (normalizedProps.text === undefined || normalizedProps.text === null || normalizedProps.text === '')) {
      normalizedProps.text = resolvedText
    }

    const normalized = {
      ...raw,
      text: resolvedText,
      props: normalizedProps
    }
    
    return normalized
  }

  if (typeof raw === 'string') {
    return {
      id: 'text',
      type: 'text',
      text: raw,
      props: { text: raw }
    }
  }

  return raw
}

/**
 * تعیین نام کامپوننت برای رندر ویجت فوتر بر اساس فیلد type یا component
 */
function resolveWidgetKey(item: any): string | null {
  let raw: string | undefined
  if (typeof item === 'string') {
    raw = item
  } else if (item && typeof item === 'object') {
    raw = item.component || item.type || item.id || item.props?.type
  }

  if (!raw) return null

  const cleaned = raw.replace(/\.(vue|ts|js)$/gi, '')
  const segments = cleaned.split('/')
  const baseName = segments[segments.length - 1] || cleaned
  const normalizedBase = baseName.replace(/^(footerwidget|footer-widget)/i, '') || baseName
  const pascal = toPascalCase(normalizedBase)

  const candidates = new Set<string>([
    baseName,
    cleaned,
    pascal,
    `FooterWidget${pascal}`
  ])

  // اگر نام اصلی خودش با FooterWidget شروع می‌شود، همان را هم اضافه کن
  if (/^footerwidget/i.test(baseName)) {
    candidates.add(baseName)
  }

  for (const candidate of candidates) {
    if (candidate && footerWidgetRegistry[candidate]) {
      return candidate
    }
  }

  return null
}

function resolveWidgetComponent(item: any): string | null {
  const key = resolveWidgetKey(item)
  return key ? footerWidgetRegistry[key] ?? null : null
}

function isSocialWidget(item: any): boolean {
  const key = resolveWidgetKey(item)
  if (!key) return false
  const isSocial = key.toLowerCase() === 'footerwidgetsocial'
  
  return isSocial
}

function isContactWidget(item: any): boolean {
  const key = resolveWidgetKey(item)
  if (!key) {
    return item?.id === 'contact'
  }
  return key.toLowerCase() === 'footerwidgetcontact'
}

function getWidgetProps(item: any) {
  const baseProps: Record<string, any> = {}

  if (item.paddingRight !== undefined) baseProps.paddingRight = item.paddingRight
  if (item.paddingLeft !== undefined) baseProps.paddingLeft = item.paddingLeft
  if (item.paddingTop !== undefined) baseProps.paddingTop = item.paddingTop
  if (item.paddingBottom !== undefined) baseProps.paddingBottom = item.paddingBottom
  if (item.align !== undefined) baseProps.align = item.align

  // Props مخصوص تصویر
  if (item.id === 'image') {
    baseProps.imageUrl = item.imageUrl
    baseProps.imageName = item.imageName
    baseProps.imageId = item.imageId
  }

  if (isSocialWidget(item)) {
    const explicitSocials = normalizeSocialEntries(item.props?.socials ?? item.socials)
    const globalSocials = normalizeSocialEntries(socialMediaSettings.value?.socials)
    const socials = explicitSocials.length > 0 ? explicitSocials : globalSocials
    baseProps.socials = socials
  }

  // Props مخصوص اطلاعات تماس
  if (isContactWidget(item) && companySettings.value) {
    // Map location-aware settings into the legacy props the widget expects
    const settings = companySettings.value
    const locationList = Array.isArray(settings.locations) ? settings.locations : []
    const locationAddresses = locationList
      .map((location: any) => typeof location?.address === 'string' ? normalizeText(location.address) : '')
      .filter(Boolean)

    const fallbackAddress = typeof settings.address === 'string' ? normalizeText(settings.address) : ''
    const fallbackSecondaryAddress = typeof settings.secondaryAddress === 'string' ? normalizeText(settings.secondaryAddress) : ''

    baseProps.address = locationAddresses[0] || fallbackAddress

    const secondaryFromLocations = locationAddresses.slice(1).find(Boolean) || ''
    baseProps.secondaryAddress = secondaryFromLocations || (fallbackSecondaryAddress && fallbackSecondaryAddress !== baseProps.address ? fallbackSecondaryAddress : '')

    const basePhones = Array.isArray(settings.phones) ? settings.phones : parsePhoneList(settings.phones)
    const locationPhones = locationList.flatMap((location: any) => Array.isArray(location?.phones) ? location.phones : [])
    baseProps.phones = Array.from(new Set([...(basePhones || []), ...locationPhones].map((phone: string) => normalizePhoneValue(phone)))).filter(Boolean)

    const baseMobiles = Array.isArray(settings.mobiles) ? settings.mobiles : parsePhoneList(settings.mobiles)
    const locationMobiles = locationPhones.filter((phone: string) => isLikelyMobileNumber(phone))
    baseProps.mobiles = Array.from(new Set([...(baseMobiles || []), ...locationMobiles].map((phone: string) => normalizePhoneValue(phone)))).filter(Boolean)

    let email = typeof settings.email === 'string' ? settings.email : ''
    if (!email) {
      const locationWithEmail = locationList.find((location: any) => typeof location?.email === 'string' && location.email.trim())
      if (locationWithEmail?.email) {
        email = locationWithEmail.email
      }
    }

    if (locationList.length > 0) {
      baseProps.locations = locationList
    }

    baseProps.email = email || ''
    baseProps.title = item.props?.title || ''
  }

  if (item.id === 'about') {
    if (companySettings.value) {
      baseProps.description = item.props?.description ?? (companySettings.value.about || '')
      baseProps.title = item.props?.title || 'درباره ما'
    }
  }

  // Props مخصوص ساعات کاری - از تنظیمات شرکت
  if (item.id === 'working-hours') {
    if (companySettings.value) {
  baseProps.workingSchedules = companySettings.value.workingSchedules || []
  baseProps.title = item.props?.title || 'ساعات کاری'
    }
  }

  // اضافه کردن بقیه props که مشخص نشده‌اند
  const allProps = {
    ...(typeof item === 'object' ? (item.props || {}) : {}),
    ...baseProps
  }

  return allProps
}

function toPascalCase(str: string): string {
  return str
    .replace(/[^a-zA-Z0-9]+/g, ' ') // جداکننده‌ها را فاصله می‌کنیم
    .split(' ')
    .map(segment => segment.charAt(0).toUpperCase() + segment.slice(1))
    .join('')
}

// تابع تعیین justify-content بر اساس چینش
function getJustifyContent(align: string): string {
  switch (align) {
    case 'left':
      return 'flex-start'  // در RTL: چپ = flex-start
    case 'center':
      return 'center'
    case 'right':
      return 'flex-end'  // در RTL: راست = flex-end
    default:
      return 'center'
  }
}

function getItemStyle(item: any) {
  if (typeof item !== 'object') return {}
  const style: Record<string, any> = {}

  // تنظیم عرض آیتم
  if (typeof item.width === 'number' && item.width > 0) {
    style.flex = `0 0 ${item.width}%`
    style.width = `${item.width}%`
  }

  if (item.height) {
    style.minHeight = `${item.height}px`
  }

  // تنظیم پس‌زمینه
  if (item.bgColor && item.bgColor !== 'transparent') {
    style.backgroundColor = item.bgColor
  }

  // تنظیم پدینگ راست و چپ
  if (item.paddingRight !== undefined) {
    style.paddingRight = `${item.paddingRight}px`
  }
  if (item.paddingLeft !== undefined) {
    style.paddingLeft = `${item.paddingLeft}px`
  }

  // تنظیم چینش آیتم
  if (item.align) {
    style.justifyContent = getJustifyContent(item.align)
  }

  return style
}

/**
 * تعیین استایل جداکننده لایه
 */
function getSeparatorStyle(layer: any) {
  const style: Record<string, any> = {}

  // تنظیم نوع خط
  if (layer.separatorType) {
    switch (layer.separatorType) {
      case 'dashed':
        style.borderLeftStyle = 'dashed'
        break
      case 'dotted':
        style.borderLeftStyle = 'dotted'
        break
      case 'double':
        style.borderLeftStyle = 'double'
        break
      default:
        style.borderLeftStyle = 'solid'
    }
  } else {
    style.borderLeftStyle = 'solid'
  }

  // رنگ جداکننده با شفافیت
  const separatorColor = layer.separatorColor || '#ffffff'
  const separatorOpacity = layer.separatorOpacity !== undefined ? layer.separatorOpacity / 100 : 0.2
  
  // تبدیل رنگ hex به rgba
  let rgba = 'rgba(255, 255, 255, 0.2)'
  if (separatorColor.startsWith('#')) {
    const hex = separatorColor.replace('#', '')
    const r = parseInt(hex.substring(0, 2), 16)
    const g = parseInt(hex.substring(2, 4), 16)
    const b = parseInt(hex.substring(4, 6), 16)
    rgba = `rgba(${r}, ${g}, ${b}, ${separatorOpacity})`
  } else if (separatorColor.startsWith('rgb')) {
    // اگر از قبل rgba یا rgb است
    rgba = separatorColor.replace('rgb(', 'rgba(').replace(')', `, ${separatorOpacity})`)
  }
  
  style.borderLeftColor = rgba
  style.borderLeftWidth = `${layer.separatorWidth || 1}px`
  style.height = '60%'

  return style
}

// @ts-ignore - vite/nuxt import meta glob
const footerWidgetModules = import.meta.glob('./FooterWidgets/*.vue', { eager: true })
const footerWidgetRegistry: Record<string, any> = {}
for (const path in footerWidgetModules) {
  const mod: any = (footerWidgetModules as any)[path]
  if (mod && mod.default) {
    const name = path.split('/').pop()?.replace('.vue', '') || ''
    footerWidgetRegistry[name] = mod.default
  }
}

// Watch برای تغییرات route (بیرون از onMounted)
watch(() => route?.path, async (newPath) => {
  if (!newPath) return
  
  // بررسی مجدد اینکه آیا فوتر باید نمایش داده شود
  if (activeFooter.value) {
    try {
      const footerForPage = getFooterForPage(newPath)

      // اگر فوتر برای صفحه جدید وجود نداشته باشد، فوتر را پاک کن
      if (!footerForPage) {
        activeFooter.value = null
      } else {
        // اگر فوتر برای صفحه جدید وجود داشته باشد، آن را تنظیم کن
        activeFooter.value = footerForPage
      }
    } catch {
      // در صورت خطا، فوتر را پاک کن
      activeFooter.value = null
    }
  }
})

onMounted(async () => {
  try {
    // اجرای همزمان درخواست‌های داده‌ای فوتر برای کاهش زمان انتظار
    await Promise.all([
      getSocialMediaSettings(),
      getCompanySettings(),
      loadFooters()
    ])
    activeFooter.value = getActiveFooter()
    
    // Parse styleSettings برای هر لایه
    if (activeFooter.value && activeFooter.value.layers) {
      activeFooter.value.layers = activeFooter.value.layers.map((layer: any) => {
        let styleSettings: any = { border: {}, shadow: {}, layout: {} }
        if (layer.styleSettings) {
          try {
            styleSettings = typeof layer.styleSettings === 'string' 
              ? JSON.parse(layer.styleSettings) 
              : layer.styleSettings
          } catch {
            styleSettings = { border: {}, shadow: {}, layout: {} }
          }
        }
        
        return {
          ...layer,
          enableBorder: styleSettings.border?.enabled || false,
          borderPosition: styleSettings.border?.position || 'all',
          borderColor: styleSettings.border?.color || '#e5e7eb',
          borderWidth: styleSettings.border?.width || 1,
          borderStyle: styleSettings.border?.style || 'solid',
          enableShadow: styleSettings.shadow?.enabled || false,
          shadowIntensity: styleSettings.shadow?.intensity || 'md',
          shadowDirection: styleSettings.shadow?.direction || 'top'
        }
      })
    }
  } catch {
    activeFooter.value = null
  }
})
</script>

<style scoped>
/* Prevent horizontal scrolling */
footer {
  width: 100%;
  max-width: 100vw;
  overflow: hidden;
}

.footer-layer {
  display: flex;
  align-items: center;
}

.footer-item-wrapper {
  width: 100%;
  max-width: 100%;
  overflow: hidden;
  box-sizing: border-box;
}

.footer-separator {
  align-self: stretch;
  margin: 0 8px;
  border-left: 1px solid rgba(255, 255, 255, 0.2);
  height: auto;
  min-height: 20px;
}

/* استایل‌های responsive */
@media (max-width: 768px) {
  .footer-item-wrapper {
    min-height: 40px;
    width: 100%;
    max-width: 100vw;
    overflow-x: hidden;
  }
  
  .footer-separator {
    margin: 0 4px;
  }
}

@media (max-width: 640px) {
  footer {
    width: 100vw;
    margin-left: calc(-50vw + 50%);
    margin-right: calc(-50vw + 50%);
  }
}
</style>
