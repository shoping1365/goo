// Temporarily hide heavy backup assets from Nuxt/Nitro public-assets scanning during build on Windows
// Renames `public/uploads/media/backup` -> `public/uploads/media/.backup_hidden`

const fs = require('fs');
const fsp = require('fs/promises');
const path = require('path');

async function pathExists(p) {
     try {
          await fsp.stat(p);
          return true;
     } catch {
          return false;
     }
}

async function main() {
     const projectRoot = process.cwd();
     const backupDir = path.join(projectRoot, 'public', 'uploads', 'media', 'backup');
     const hiddenDir = path.join(projectRoot, 'public', 'uploads', 'media', '.backup_hidden');

     try {
          const exists = await pathExists(backupDir);
          if (!exists) {
               console.log('[prebuild] No backup directory to hide.');
               return;
          }

          const hiddenExists = await pathExists(hiddenDir);
          if (hiddenExists) {
               console.log('[prebuild] Hidden backup directory already exists. Skipping.');
               return;
          }

          await fsp.rename(backupDir, hiddenDir);
          console.log('[prebuild] Backup directory temporarily hidden:', hiddenDir);
     } catch (err) {
          console.warn('[prebuild] Warning: could not hide backup directory:', err?.message || err);
          // Do not fail the build due to a non-critical step
     }
}

main();


