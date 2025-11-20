import { extname } from 'path'
import { promisify } from 'util'
import { Plugin } from 'vite'
import { createBrotliCompress, createGzip } from 'zlib'

const gzip = promisify(createGzip)
const brotli = promisify(createBrotliCompress)

/**
 * پلاگین فشرده‌سازی برای Vite
 * فایل‌های CSS و JS را با Gzip و Brotli فشرده می‌کند
 */
export function compressionPlugin(): Plugin {
  return {
    name: 'compression-plugin',
    apply: 'build', // فقط در build فعال می‌شود
    async generateBundle(options, bundle) {
      // فقط در production فعال می‌شود
      if (process.env.NODE_ENV !== 'production') {
        return
      }

      const filesToCompress = Object.keys(bundle).filter(fileName => {
        const ext = extname(fileName)
        return ext === '.js' || ext === '.css'
      })

      for (const fileName of filesToCompress) {
        const chunk = bundle[fileName]
        if (chunk.type === 'chunk' || chunk.type === 'asset') {
          const source = chunk.type === 'chunk' ? chunk.code : chunk.source
          
          try {
            // فشرده‌سازی Gzip
            const gzipped = await gzip(Buffer.from(source))
            this.emitFile({
              type: 'asset',
              fileName: `${fileName}.gz`,
              source: gzipped
            })

            // فشرده‌سازی Brotli
            const brotlied = await brotli(Buffer.from(source))
            this.emitFile({
              type: 'asset',
              fileName: `${fileName}.br`,
              source: brotlied
            })

            // console.log(`✅ فشرده‌سازی ${fileName}: Gzip ${gzipped.length} bytes, Brotli ${brotlied.length} bytes`)
          } catch (error) {
            console.error(`❌ خطا در فشرده‌سازی ${fileName}:`, error)
          }
        }
      }
    }
  }
}








































