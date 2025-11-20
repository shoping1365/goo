<template>
  <span>{{ formattedDate }}</span>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  dateString: {
    type: String,
    required: true
  },
  showTime: {
    type: Boolean,
    default: false
  },
  format: {
    type: String,
    default: 'YYYY/MM/DD', // YYYY/MM/DD, HH:MM, HH:MM - YYYY/MM/DD
    validator: (value) => ['YYYY/MM/DD', 'HH:MM', 'HH:MM - YYYY/MM/DD', 'YYYY/MM/DD - HH:MM'].includes(value)
  }
})

// تابع تبدیل میلادی به شمسی
function gregorianToJalali(gy, gm, gd) {
  const g_d_m = [0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334]
  let jy = (gy <= 1600) ? 0 : 979
  gy -= (gy <= 1600) ? 621 : 1600
  const gy2 = (gm > 2) ? (gy + 1) : gy
  let days = (365 * gy) + (parseInt((gy2 + 3) / 4)) - (parseInt((gy2 + 99) / 100)) + (parseInt((gy2 + 399) / 400)) - 80 + gd + g_d_m[gm - 1]
  jy += 33 * (parseInt(days / 12053))
  days %= 12053
  jy += 4 * (parseInt(days / 1461))
  days %= 1461
  jy += parseInt((days - 1) / 365)
  if (days > 365) days = (days - 1) % 365
  const jm = (days < 186) ? 1 + parseInt(days / 31) : 7 + parseInt((days - 186) / 30)
  const jd = 1 + ((days < 186) ? (days % 31) : ((days - 186) % 30))
  return { jy: jy, jm: jm, jd: jd }
}

const formattedDate = computed(() => {
  if (!props.dateString) return 'نامشخص'
  
  try {
    const date = new Date(props.dateString)
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
    switch (props.format) {
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
})
</script>

