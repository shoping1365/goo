const fs = require('fs');
const path = require('path');

// مسیرهای فایل‌های ویجت
const widgetPaths = [
  'app/pages/admin/content/banners/pentabanner/create/index.vue',
  'app/pages/admin/content/banners/pentabanner/edit/[id].vue',
  'app/pages/admin/content/banners/FullBanner/create/index.vue',
  'app/pages/admin/content/banners/FullBanner/edit/[id].vue',
  'app/pages/admin/content/banners/DoubleBanner/create/index.vue',
  'app/pages/admin/content/banners/DoubleBanner/edit/[id].vue',
  'app/pages/admin/content/banners/triplebanner/create/index.vue',
  'app/pages/admin/content/banners/triplebanner/edit/[id].vue',
  'app/pages/admin/content/banners/quadbanner/create/index.vue',
  'app/pages/admin/content/banners/quadbanner/edit/[id].vue'
];

function updateWidgetFile(filePath) {
  try {
    const fullPath = path.join(process.cwd(), filePath);
    
    if (!fs.existsSync(fullPath)) {
      console.log(`فایل یافت نشد: ${filePath}`);
      return;
    }

    let content = fs.readFileSync(fullPath, 'utf8');
    
    // اضافه کردن useToast import
    if (!content.includes('useToast')) {
      // پیدا کردن خط composables
      const composablesMatch = content.match(/(\/\/ Composables[\s\S]*?const.*?useWidget\(\))/);
      if (composablesMatch) {
        const replacement = composablesMatch[1] + '\nconst { showSuccess, showError } = useToast()';
        content = content.replace(composablesMatch[1], replacement);
      }
    }
    
    // حذف navigateTo و اضافه کردن showSuccess
    if (content.includes('await navigateTo(\'/admin/content/banners\')')) {
      content = content.replace(
        /\/\/ Redirect to widgets list page after successful creation[\s\S]*?await navigateTo\('\/admin\/content\/banners'\)/,
        "showSuccess('ابزارک با موفقیت ایجاد شد!')"
      );
      
      content = content.replace(
        /\/\/ Redirect to widgets list page after successful update[\s\S]*?await navigateTo\('\/admin\/content\/banners'\)/,
        "showSuccess('ابزارک با موفقیت به\u200cروزرسانی شد!')"
      );
    }
    
    fs.writeFileSync(fullPath, content, 'utf8');
    console.log(`✅ به‌روزرسانی شد: ${filePath}`);
    
  } catch (error) {
    console.error(`❌ خطا در به‌روزرسانی ${filePath}:`, error.message);
  }
}

// اجرای اسکریپت
console.log('🚀 شروع به‌روزرسانی فایل‌های ویجت...\n');

widgetPaths.forEach(updateWidgetFile);

console.log('\n✨ تمام فایل‌ها به‌روزرسانی شدند!');
