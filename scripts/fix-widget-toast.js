#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// لیست فایل‌های ویجت
const files = [
  'app/pages/admin/content/banners/pentabanner/edit/[id].vue',
  'app/pages/admin/content/banners/FullBanner/create/index.vue',
  'app/pages/admin/content/banners/FullBanner/edit/[id].vue',
  'app/pages/admin/content/banners/DoubleBanner/create/index.vue',
  'app/pages/admin/content/banners/DoubleBanner/edit/[id].vue',
  'app/pages/admin/content/banners/triplebanner/create/index.vue',
  'app/pages/admin/content/banners/triplebanner/edit/[id].vue'
];

files.forEach(filePath => {
  const fullPath = path.join(process.cwd(), filePath);
  
  if (!fs.existsSync(fullPath)) {
    console.log(`❌ فایل یافت نشد: ${filePath}`);
    return;
  }

  let content = fs.readFileSync(fullPath, 'utf8');
  
  // اضافه کردن useToast
  if (!content.includes('useToast')) {
    content = content.replace(
      /(\/\/ Composables[\s\S]*?const.*?useWidget\(\))/,
      '$1\nconst { showSuccess, showError } = useToast()'
    );
  }
  
  // جایگزینی navigateTo با showSuccess
  content = content.replace(
    /\/\/ Redirect to widgets list page after successful creation[\s\S]*?await navigateTo\('\/admin\/content\/banners'\)/,
    "showSuccess('ابزارک با موفقیت ایجاد شد!')"
  );
  
  content = content.replace(
    /\/\/ Redirect to widgets list page after successful update[\s\S]*?await navigateTo\('\/admin\/content\/banners'\)/,
    "showSuccess('ابزارک با موفقیت به\u200cروزرسانی شد!')"
  );
  
  fs.writeFileSync(fullPath, content, 'utf8');
  console.log(`✅ به‌روزرسانی شد: ${filePath}`);
});

console.log('\n🎉 تمام فایل‌ها به‌روزرسانی شدند!');
