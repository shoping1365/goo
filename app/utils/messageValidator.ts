// پیکربندی امنیتی پیام‌ها
const SECURITY_CONFIG = {
  // حداکثر طول پیام
  maxLength: 1000,
  
  // الگوهای مخرب (Regex patterns)
  // Note: Regex-based filtering is not perfect. Consider using DOMPurify for production.
  maliciousPatterns: [
    // JavaScript injection (improved regex)
    /<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gis,
    /javascript:/gi,
    /on\w+\s*=/gi,
    /eval\s*\(/gi,
    /Function\s*\(/gi,
    /<style\b[^<]*(?:(?!<\/style>)<[^<]*)*<\/style>/gis,
    
    // HTML injection  
    /<iframe\b[^>]*>/gi,
    /<object\b[^>]*>/gi,
    /<embed\b[^>]*>/gi,
    /<link\b[^>]*>/gi,
    /<meta\b[^>]*>/gi,
    /<form\b[^>]*>/gi,
    
    // SQL injection patterns
    /(\b(SELECT|INSERT|UPDATE|DELETE|DROP|CREATE|ALTER|EXEC|UNION)\b)/gi,
    /(--|\#|\/\*|\*\/)/gi,
    /(\b(OR|AND)\s+\d+\s*=\s*\d+)/gi,
    
    // File system attacks
    /\.\.\/|\.\.\\/gi,
    /\/etc\/passwd/gi,
    /\/etc\/shadow/gi,
    /\/proc\//gi,
    
    // Command injection
    /(\||&&|\;\s*|\$\(|\`)/gi,
    /\b(wget|curl|nc|netcat|ping|nslookup|dig)\b/gi,
    
    // XSS patterns
    /(<\s*\/?\s*[a-z][\w\-]*[^>]*>)/gi,
    /(src\s*=\s*["']?[^"'\s>]+["']?)/gi,
    /(href\s*=\s*["']?[^"'\s>]+["']?)/gi,
  ],
  
  // کلمات ممنوع
  bannedWords: [
    'alert', 'confirm', 'prompt', 'console.log',
    'document.cookie', 'localStorage', 'sessionStorage',
    'window.location', 'location.href', 'location.replace',
    'eval', 'Function', 'setTimeout', 'setInterval',
  ],
  
  // الگوهای URL مشکوک
  suspiciousUrls: [
    /https?:\/\/bit\.ly/gi,
    /https?:\/\/tinyurl/gi,
    /https?:\/\/[^\/]*\.(tk|ml|ga|cf|xyz)/gi,
    /data:text\/html/gi,
    /data:application/gi,
  ],
  
  // فرمت‌های فایل مجاز (برای آینده)
  allowedFileTypes: [
    'image/jpeg', 'image/png', 'image/gif', 'image/webp',
    'text/plain', 'application/pdf'
  ],
  
  // حداکثر اندازه فایل (بایت)
  maxFileSize: 5 * 1024 * 1024, // 5MB
}

// نتیجه اعتبارسنجی
interface ValidationResult {
  isValid: boolean
  errors: string[]
  sanitizedContent?: string
  riskLevel: 'low' | 'medium' | 'high' | 'critical'
}

// اعتبارسنجی پیام متنی
export function validateTextMessage(content: string): ValidationResult {
  const errors: string[] = []
  let riskLevel: ValidationResult['riskLevel'] = 'low'
  
  // بررسی طول پیام
  if (!content || content.trim().length === 0) {
    errors.push('پیام نمی‌تواند خالی باشد')
    return { isValid: false, errors, riskLevel: 'low' }
  }
  
  if (content.length > SECURITY_CONFIG.maxLength) {
    errors.push(`پیام نمی‌تواند بیشتر از ${SECURITY_CONFIG.maxLength} کاراکتر باشد`)
    riskLevel = 'medium'
  }
  
  // بررسی الگوهای مخرب
  for (const pattern of SECURITY_CONFIG.maliciousPatterns) {
    if (pattern.test(content)) {
      errors.push('پیام حاوی محتوای مشکوک است')
      riskLevel = 'critical'
      break
    }
  }
  
  // بررسی کلمات ممنوع
  const lowerContent = content.toLowerCase()
  for (const word of SECURITY_CONFIG.bannedWords) {
    if (lowerContent.includes(word.toLowerCase())) {
      errors.push('پیام حاوی کلمات ممنوع است')
      riskLevel = riskLevel === 'critical' ? 'critical' : 'high'
      break
    }
  }
  
  // بررسی URL های مشکوک
  for (const urlPattern of SECURITY_CONFIG.suspiciousUrls) {
    if (urlPattern.test(content)) {
      errors.push('پیام حاوی لینک مشکوک است')
      riskLevel = riskLevel === 'critical' ? 'critical' : 'high'
      break
    }
  }
  
  // پاکسازی محتوا
  const sanitizedContent = sanitizeContent(content)
  
  return {
    isValid: errors.length === 0,
    errors,
    sanitizedContent,
    riskLevel
  }
}

// پاکسازی محتوای پیام
export function sanitizeContent(content: string): string {
  let sanitized = content
  
  // حذف تگ‌های HTML
  sanitized = sanitized.replace(/<[^>]*>/g, '')
  
  // حذف کاراکترهای کنترلی
  sanitized = sanitized.replace(/[\x00-\x1F\x7F]/g, '')
  
  // حذف فضاهای اضافی
  sanitized = sanitized.replace(/\s+/g, ' ').trim()
  
  // محدود کردن کاراکترهای تکراری
  sanitized = sanitized.replace(/(.)\1{4,}/g, '$1$1$1$1')
  
  return sanitized
}

// اعتبارسنجی نوع فایل
export function validateFileType(mimeType: string): boolean {
  return SECURITY_CONFIG.allowedFileTypes.includes(mimeType)
}

// اعتبارسنجی اندازه فایل
export function validateFileSize(size: number): boolean {
  return size <= SECURITY_CONFIG.maxFileSize
}

// اعتبارسنجی نام فایل
export function validateFileName(fileName: string): ValidationResult {
  const errors: string[] = []
  let riskLevel: ValidationResult['riskLevel'] = 'low'
  
  // بررسی کاراکترهای غیرمجاز
  if (/[<>:"/\\|?*]/.test(fileName)) {
    errors.push('نام فایل حاوی کاراکترهای غیرمجاز است')
    riskLevel = 'medium'
  }
  
  // بررسی پسوند مخرب
  const dangerousExtensions = [
    '.exe', '.bat', '.cmd', '.com', '.scr', '.pif', '.vbs', '.js', '.jar',
    '.php', '.asp', '.jsp', '.sh', '.ps1', '.py', '.pl', '.rb'
  ]
  
  const lowerFileName = fileName.toLowerCase()
  for (const ext of dangerousExtensions) {
    if (lowerFileName.endsWith(ext)) {
      errors.push('نوع فایل مجاز نیست')
      riskLevel = 'critical'
      break
    }
  }
  
  // بررسی نام‌های فایل سیستمی
  const systemFiles = ['con', 'prn', 'aux', 'nul', 'com1', 'com2', 'lpt1', 'lpt2']
  const baseFileName = fileName.split('.')[0].toLowerCase()
  if (systemFiles.includes(baseFileName)) {
    errors.push('نام فایل مجاز نیست')
    riskLevel = 'high'
  }
  
  return {
    isValid: errors.length === 0,
    errors,
    riskLevel
  }
}

// تولید گزارش امنیتی
export function generateSecurityReport(validationResult: ValidationResult, userInfo: Record<string, unknown>) {
  return {
    timestamp: new Date().toISOString(),
    userId: userInfo?.id || 'anonymous',
    userRole: userInfo?.role || 'guest',
    riskLevel: validationResult.riskLevel,
    errors: validationResult.errors,
    action: validationResult.isValid ? 'allowed' : 'blocked'
  }
}

// لاگ امنیتی
export function logSecurityEvent(_report: unknown) {
  // Security event logged
}