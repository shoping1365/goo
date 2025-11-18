import { defineNuxtPlugin } from 'nuxt/app'

// پلاگین مدیریت فونت‌های فارسی برای Nuxt 4
export default defineNuxtPlugin(() => {
  // اضافه کردن CSS فونت‌ها به head
  // Preload فونت‌های اصلی
  const preloadFonts = [
    '/fonts/IRANSansWeb.woff2',
    '/fonts/Vazir.woff2',
    '/fonts/Shabnam.woff2'
  ]

  preloadFonts.forEach(href => {
    const link = document.createElement('link')
    link.rel = 'preload'
    link.as = 'font'
    link.type = 'font/woff2'
    link.crossOrigin = 'anonymous'
    link.href = href
    document.head.appendChild(link)
  })

  // اضافه کردن CSS برای فونت‌های فارسی
  const style = document.createElement('style')
  style.textContent = `
    /* فونت‌های فارسی - بهینه‌سازی شده برای Nuxt 4 */
    
    /* IRANSansWeb - فونت اصلی */
    @font-face {
      font-family: 'IRANSansWeb';
      src: url('/fonts/IRANSansWeb.woff2') format('woff2'),
      src: url('/fonts/IRANSansWeb.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWeb';
      src: url('/fonts/IRANSansWeb_Bold.woff2') format('woff2'),
      src:url('/fonts/IRANSansWeb_Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWeb';
      src: url('/fonts/IRANSansWeb_Light.woff2') format('woff2'),
      src: url('/fonts/IRANSansWeb_Light.woff') format('woff');
      font-weight: 300;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWeb';
      src: url('/fonts/IRANSansWeb_Medium.woff2') format('woff2'),
      url('/fonts/IRANSansWeb_Medium.woff') format('woff');
      font-weight: 500;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWeb';
      src: url('/fonts/IRANSansWeb_Black.woff2') format('woff2'),
      url('/fonts/IRANSansWeb_Black.woff') format('woff');
      font-weight: 900;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWeb';
      src: url('/fonts/IRANSansWeb_UltraLight.woff2') format('woff2'),
      url('/fonts/IRANSansWeb_UltraLight.woff') format('woff');
      font-weight: 200;
      font-style: normal;
      font-display: swap;
    }
    
    /* IRANSansWebFaNum - با اعداد فارسی */
    @font-face {
      font-family: 'IRANSansWebFaNum';
      src: url('/fonts/IRANSansWebFaNum.woff2') format('woff2'),
      url('/fonts/IRANSansWebFaNum.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWebFaNum';
      src: url('/fonts/IRANSansWebFaNum_Bold.woff2') format('woff2'),
      url('/fonts/IRANSansWebFaNum_Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWebFaNum';
      src: url('/fonts/IRANSansWebFaNum_Light.woff2') format('woff2'),
      src: url('/fonts/IRANSansWebFaNum_Light.woff') format('woff');
      font-weight: 300;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWebFaNum';
      src: url('/fonts/IRANSansWebFaNum_Medium.woff2') format('woff2'),
      src: url('/fonts/IRANSansWebFaNum_Medium.woff') format('woff');
      font-weight: 500;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWebFaNum';
      src: url('/fonts/IRANSansWebFaNum_Black.woff2') format('woff2'),
      src: url('/fonts/IRANSansWebFaNum_Black.woff') format('woff');
      font-weight: 900;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'IRANSansWebFaNum';
      src: url('/fonts/IRANSansWebFaNum_UltraLight.woff2') format('woff2'),
      src: url('/fonts/IRANSansWebFaNum_UltraLight.woff') format('woff');
      font-weight: 200;
      font-style: normal;
      font-display: swap;
    }
    
    /* Vazir - فونت وزیر */
    @font-face {
      font-family: 'Vazir';
      src: url('/fonts/Vazir.woff2') format('woff2'),
      src: url('/fonts/Vazir.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Vazir';
      src: url('/fonts/Vazir-Bold.woff2') format('woff2'),
      src: url('/fonts/Vazir-Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Vazir';
      src: url('/fonts/Vazir-Light.woff2') format('woff2'),
      src: url('/fonts/Vazir-Light.woff') format('woff');
      font-weight: 300;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Vazir';
      src: url('/fonts/Vazir-Medium.woff2') format('woff2'),
      src: url('/fonts/Vazir-Medium.woff') format('woff');
      font-weight: 500;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Vazir';
      src: url('/fonts/Vazir-Thin.woff2') format('woff2'),
      src: url('/fonts/Vazir-Thin.woff') format('woff');
      font-weight: 100;
      font-style: normal;
      font-display: swap;
    }
    
    /* Shabnam - فونت شبنم */
    @font-face {
      font-family: 'Shabnam';
      src: url('/fonts/Shabnam.woff2') format('woff2'),
      src: url('/fonts/Shabnam.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Shabnam';
      src: url('/fonts/Shabnam-Bold.woff2') format('woff2'),
      src: url('/fonts/Shabnam-Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Shabnam';
      src: url('/fonts/Shabnam-Light.woff2') format('woff2'),
      src: url('/fonts/Shabnam-Light.woff') format('woff');
      font-weight: 300;
      font-style: normal;
      font-display: swap;
    }
    
    /* Sahel - فونت ساحل */
    @font-face {
      font-family: 'Sahel';
      src: url('/fonts/Sahel.woff2') format('woff2'),
      src: url('/fonts/Sahel.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Sahel';
      src: url('/fonts/Sahel-Bold.woff2') format('woff2'),
      src: url('/fonts/Sahel-Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Sahel';
      src: url('/fonts/Sahel-Black.woff2') format('woff2'),
      src: url('/fonts/Sahel-Black.woff') format('woff');
      font-weight: 900;
      font-style: normal;
      font-display: swap;
    }
    
    /* Samim - فونت سمیم */
    @font-face {
      font-family: 'Samim';
      src: url('/fonts/Samim-Bold.woff') format('woff'),
      src: url('/fonts/Samim-Bold.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Samim';
      src: url('/fonts/Samim-Bold.woff2') format('woff2'),
      src: url('/fonts/Samim-Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    /* Parastoo - فونت پرستو */
    @font-face {
      font-family: 'Parastoo';
      src: url('/fonts/Parastoo.woff2') format('woff2'),
      src: url('/fonts/Parastoo.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    @font-face {
      font-family: 'Parastoo';
      src: url('/fonts/Parastoo-Bold.woff2') format('woff2'),
      src: url('/fonts/Parastoo-Bold.woff') format('woff');
      font-weight: 700;
      font-style: normal;
      font-display: swap;
    }
    
    /* Tanha - فونت تنها */
    @font-face {
      font-family: 'Tanha';
      src: url('/fonts/Tanha.woff2') format('woff2'),
      src: url('/fonts/Tanha.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    /* Yekan - فونت یکان */
    @font-face {
      font-family: 'Yekan';
      src: url('/fonts/yekan-regular.woff') format('woff');
      font-weight: 400;
      font-style: normal;
      font-display: swap;
    }
    
    /* تنظیم فونت پیش‌فرض برای RTL */
    [dir="rtl"], .rtl {
      font-family: 'IRANSansWeb', 'IRANSansWebFaNum', 'Vazir', 'Shabnam', Tahoma, Arial, sans-serif;
    }
    
    /* کلاس‌های کمکی برای فونت‌ها */
    .font-iransans { font-family: 'IRANSansWeb', sans-serif; }
    .font-iransansfanum { font-family: 'IRANSansWebFaNum', sans-serif; }
    .font-vazir { font-family: 'Vazir', sans-serif; }
    .font-shabnam { font-family: 'Shabnam', sans-serif; }
    .font-sahel { font-family: 'Sahel', sans-serif; }
    .font-samim { font-family: 'Samim', sans-serif; }
    .font-parastoo { font-family: 'Parastoo', sans-serif; }
    .font-tanha { font-family: 'Tanha', sans-serif; }
    .font-yekan { font-family: 'Yekan', sans-serif; }
  `
  document.head.appendChild(style)
}) 