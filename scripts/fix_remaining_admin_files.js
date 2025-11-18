const fs = require('fs');
const path = require('path');

// لیست فایل‌های باقی مانده که نیاز به اصلاح دارند
const filesToFix = [
     'app/pages/admin/content/banners/Categories/create/index.vue',
     'app/pages/admin/content/banners/DoubleBanner/index.vue',
     'app/pages/admin/content/banners/DoubleBanner/create/index.vue',
     // اضافه کردن بقیه فایل‌ها...
];

function fixFile(filePath) {
     try {
          let content = fs.readFileSync(filePath, 'utf8');

          // اگر فایل قبلاً useAuth دارد، رد کردن
          if (content.includes('import { useAuth }')) {
               console.log(`Skipping ${filePath} - already has useAuth`);
               return;
          }

          // پیدا کردن script setup
          const scriptSetupMatch = content.match(/(<script setup lang="ts">[\s\S]*?definePageMeta[\s\S]*?})/);
          if (!scriptSetupMatch) {
               console.log(`No script setup found in ${filePath}`);
               return;
          }

          let scriptSetup = scriptSetupMatch[1];

          // اضافه کردن import useAuth
          if (!scriptSetup.includes('import { useAuth }')) {
               scriptSetup = scriptSetup.replace(
                    '<script setup lang="ts">',
                    '<script setup lang="ts">\nimport { useAuth } from \'~/composables/useAuth\'\n'
               );
          }

          // اضافه کردن استفاده از useAuth
          if (!scriptSetup.includes('const { user, hasPermission } = useAuth()')) {
               scriptSetup = scriptSetup.replace(
                    '})\n\n//',
                    '})\n\n// استفاده از useAuth برای چک کردن پرمیژن‌ها\nconst { user, hasPermission } = useAuth()\n\n//'
               );
          }

          // جایگزینی در فایل اصلی
          content = content.replace(scriptSetupMatch[1], scriptSetup);

          fs.writeFileSync(filePath, content, 'utf8');
          console.log(`Fixed ${filePath}`);

     } catch (error) {
          console.error(`Error fixing ${filePath}:`, error.message);
     }
}

console.log('Starting to fix remaining admin files...');
filesToFix.forEach(fixFile);
console.log('Done fixing files.');
