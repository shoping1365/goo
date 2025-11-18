// تابع تبدیل میلادی به شمسی
export function gregorianToJalali(gy: number, gm: number, gd: number) {
  const g_d_m = [0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334]
  let jy = (gy <= 1600) ? 0 : 979
  gy -= (gy <= 1600) ? 621 : 1600
  let gy2 = (gm > 2) ? (gy + 1) : gy
  let days = (365 * gy) + Math.floor((gy2 + 3) / 4) - Math.floor((gy2 + 99) / 100) + Math.floor((gy2 + 399) / 400) - 80 + gd + g_d_m[gm - 1]
  jy += 33 * Math.floor(days / 12053)
  days %= 12053
  jy += 4 * Math.floor(days / 1461)
  days %= 1461
  jy += Math.floor((days - 1) / 365)
  if (days > 365) days = (days - 1) % 365
  let jm = (days < 186) ? 1 + Math.floor(days / 31) : 7 + Math.floor((days - 186) / 30)
  let jd = 1 + ((days < 186) ? (days % 31) : ((days - 186) % 30))
  return { jy: jy, jm: jm, jd: jd }
}

export function usePersianDate() {
  const formatPersianDate = (dateString: string, format: 'YYYY/MM/DD' | 'HH:MM' | 'HH:MM - YYYY/MM/DD' | 'YYYY/MM/DD - HH:MM' = 'YYYY/MM/DD') => {
    if (!dateString) return 'نامشخص'
    
    try {
      const date = new Date(dateString)
      if (isNaN(date.getTime())) return 'نامشخص'
      
      // تبدیل دقیق میلادی به شمسی
      const year = date.getFullYear()
      const month = date.getMonth() + 1
      const day = date.getDate()
      const hours = date.getHours()
      const minutes = date.getMinutes()
      
      // تبدیل دقیق به شمسی
      const persianDate = gregorianToJalali(year, month, day)
      
      const yearStr = persianDate.jy.toString()
      const monthStr = persianDate.jm.toString().padStart(2, '0')
      const dayStr = persianDate.jd.toString().padStart(2, '0')
      const hoursStr = hours.toString().padStart(2, '0')
      const minutesStr = minutes.toString().padStart(2, '0')
      
      // فرمت‌بندی بر اساس format انتخابی
      switch (format) {
        case 'YYYY/MM/DD':
          return `${yearStr}/${monthStr}/${dayStr}`
        case 'HH:MM':
          return `${hoursStr}:${minutesStr}`
        case 'HH:MM - YYYY/MM/DD':
          return `${hoursStr}:${minutesStr} - ${yearStr}/${monthStr}/${dayStr}`
        case 'YYYY/MM/DD - HH:MM':
          return `${yearStr}/${monthStr}/${dayStr} - ${hoursStr}:${minutesStr}`
        default:
          return `${yearStr}/${monthStr}/${dayStr}`
      }
    } catch (error) {
      console.error('خطا در تبدیل تاریخ:', error)
      return 'نامشخص'
    }
  }

  return {
    formatPersianDate,
    gregorianToJalali
  }
}

