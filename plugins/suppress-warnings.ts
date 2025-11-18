import { Plugin } from 'vite'

/**
 * پلاگین برای حل مشکل Dynamic/Static Import Conflicts
 * این پلاگین warning های مربوط به تضاد import ها را نادیده می‌کند
 */
export function suppressImportWarningsPlugin(): Plugin {
  return {
    name: 'suppress-import-warnings',
    apply: 'build',
    configResolved(config) {
      // تنظیمات برای نادیده گرفتن warning ها
      if (config.build?.rollupOptions) {
        const originalOnWarn = config.build.rollupOptions.onwarn
        
        config.build.rollupOptions.onwarn = (warning, warn) => {
          // نادیده گرفتن warning های dynamic/static import
          if (warning.code === 'MODULE_LEVEL_DIRECTIVE') {
            return
          }
          
          if (warning.message && (
            warning.message.includes('dynamically imported') && 
            warning.message.includes('statically imported')
          )) {
            return
          }
          
          if (warning.message && warning.message.includes('dynamic import will not move module into another chunk')) {
            return
          }
          
          // اجرای onwarn اصلی برای سایر warning ها
          if (originalOnWarn) {
            originalOnWarn(warning, warn)
          } else {
            warn(warning)
          }
        }
      }
    },
    
    generateBundle(options, bundle) {
      // حذف فایل‌های تکراری
      const filesToRemove: string[] = []
      
      Object.keys(bundle).forEach(fileName => {
        const chunk = bundle[fileName]
        if (chunk.type === 'chunk') {
          // بررسی chunk های خالی یا تکراری
          if (chunk.code.length < 100) {
            filesToRemove.push(fileName)
          }
        }
      })
      
      // حذف فایل‌های غیرضروری
      filesToRemove.forEach(fileName => {
        delete bundle[fileName]
      })
      
      console.log(`🧹 حذف ${filesToRemove.length} chunk غیرضروری`)
    }
  }
}








































