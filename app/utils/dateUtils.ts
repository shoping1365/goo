// Persian Date Conversion Utilities
// ابزارهای تبدیل تاریخ میلادی به شمسی

/**
 * تبدیل تاریخ میلادی به شمسی
 * @param date - تاریخ میلادی (Date object یا string)
 * @returns تاریخ شمسی به صورت رشته
 */
export function toPersianDate(date: Date | string | null | undefined): string {
  if (!date) return '-';
  
  try {
    const dateObj = typeof date === 'string' ? new Date(date) : date;
    
    // بررسی معتبر بودن تاریخ
    if (isNaN(dateObj.getTime())) {
      return '-';
    }
    
    // تبدیل به تاریخ شمسی
    const persianDate = new Intl.DateTimeFormat('fa-IR', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      calendar: 'persian'
    }).format(dateObj);
    
    return persianDate;
  } catch (error) {
    return '-';
  }
}

/**
 * تبدیل تاریخ میلادی به شمسی با ساعت
 * @param date - تاریخ میلادی (Date object یا string)
 * @returns تاریخ و ساعت شمسی به صورت رشته (تاریخ اول، سپس ساعت)
 */
export function toPersianDateTime(date: Date | string | null | undefined): string {
  if (!date) return '-';
  
  try {
    const dateObj = typeof date === 'string' ? new Date(date) : date;
    
    // بررسی معتبر بودن تاریخ
    if (isNaN(dateObj.getTime())) {
      return '-';
    }
    
    // تبدیل تاریخ و ساعت به صورت جداگانه
    const persianDate = new Intl.DateTimeFormat('fa-IR', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      calendar: 'persian'
    }).format(dateObj);
    
    const persianTime = new Intl.DateTimeFormat('fa-IR', {
      hour: '2-digit',
      minute: '2-digit',
      calendar: 'persian'
    }).format(dateObj);
    
    // ترکیب تاریخ و ساعت با ترتیب صحیح (ساعت سمت راست)
    return `${persianTime} ${persianDate}`;
  } catch (error) {
    return '-';
  }
}

/**
 * تبدیل تاریخ میلادی به شمسی با فرمت کوتاه
 * @param date - تاریخ میلادی (Date object یا string)
 * @returns تاریخ شمسی کوتاه به صورت رشته
 */
export function toPersianDateShort(date: Date | string | null | undefined): string {
  if (!date) return '-';
  
  try {
    const dateObj = typeof date === 'string' ? new Date(date) : date;
    
    // بررسی معتبر بودن تاریخ
    if (isNaN(dateObj.getTime())) {
      return '-';
    }
    
    // تبدیل به تاریخ شمسی کوتاه
    const persianDate = new Intl.DateTimeFormat('fa-IR', {
      month: 'short',
      day: 'numeric',
      calendar: 'persian'
    }).format(dateObj);
    
    return persianDate;
  } catch (error) {
    return '-';
  }
}

/**
 * تبدیل تاریخ میلادی به شمسی با فرمت کامل
 * @param date - تاریخ میلادی (Date object یا string)
 * @returns تاریخ شمسی کامل با نام ماه
 */
export function toPersianDateFull(date: Date | string | null | undefined): string {
  if (!date) return '-';
  
  try {
    const dateObj = typeof date === 'string' ? new Date(date) : date;
    
    // بررسی معتبر بودن تاریخ
    if (isNaN(dateObj.getTime())) {
      return '-';
    }
    
    // تبدیل به تاریخ شمسی کامل
    const persianDate = new Intl.DateTimeFormat('fa-IR', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      weekday: 'long',
      calendar: 'persian'
    }).format(dateObj);
    
    return persianDate;
  } catch (error) {
    return '-';
  }
} 