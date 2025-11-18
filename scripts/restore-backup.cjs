// Restore the backup assets directory after build on Windows
// Renames `public/uploads/media/.backup_hidden` -> `public/uploads/media/backup`

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
          const hiddenExists = await pathExists(hiddenDir);
          if (!hiddenExists) {
               console.log('[postbuild] No hidden backup directory to restore.');
               return;
          }

          const backupExists = await pathExists(backupDir);
          if (backupExists) {
               console.log('[postbuild] Backup directory already present. Removing hidden copy.');
               await fsp.rm(hiddenDir, { recursive: true, force: true });
               return;
          }

          await fsp.rename(hiddenDir, backupDir);
          console.log('[postbuild] Backup directory restored:', backupDir);
     } catch (err) {
          console.warn('[postbuild] Warning: could not restore backup directory:', err?.message || err);
     }
}

main();


