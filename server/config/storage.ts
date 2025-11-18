import fs from 'fs';
import path from 'path';
import { v4 as uuidv4 } from 'uuid';

// Define storage paths
const UPLOAD_DIR = path.join(process.cwd(), 'public', 'uploads');
const MEDIA_DIR = path.join(UPLOAD_DIR, 'media');

// Create subdirectories for library
const LIBRARY_DIRS = {
  images: path.join(MEDIA_DIR, 'images', 'library'),
  videos: path.join(MEDIA_DIR, 'videos', 'library'),
  documents: path.join(MEDIA_DIR, 'documents', 'library'),
  audio: path.join(MEDIA_DIR, 'audio', 'library'),
  others: path.join(MEDIA_DIR, 'others', 'library')
};

// Ensure directories exist
function ensureDirectories() {
  if (!fs.existsSync(UPLOAD_DIR)) {
    fs.mkdirSync(UPLOAD_DIR, { recursive: true });
  }
  if (!fs.existsSync(MEDIA_DIR)) {
    fs.mkdirSync(MEDIA_DIR, { recursive: true });
  }
  
  // Create library directories
  Object.values(LIBRARY_DIRS).forEach(dir => {
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir, { recursive: true });
    }
  });
}

// Get subdirectory based on file type
function getSubdirectory(fileType: string): string {
  if (fileType.startsWith('image/')) {
    return 'images/library';
  } else if (fileType.startsWith('video/')) {
    return 'videos/library';
  } else if (fileType.startsWith('audio/')) {
    return 'audio/library';
  } else if (
    fileType.startsWith('application/pdf') ||
    fileType.startsWith('application/msword') ||
    fileType.startsWith('application/vnd.openxmlformats-officedocument') ||
    fileType.startsWith('text/')
  ) {
    return 'documents/library';
  } else {
    return 'others/library';
  }
}

// Generate unique filename
function generateUniqueFilename(originalFilename: string): string {
  const ext = path.extname(originalFilename);
  return `${uuidv4()}${ext}`;
}

// Save file to local storage
export async function saveFile(file: { filename: string; data: Buffer; type: string }) {
  ensureDirectories();
  
  const subdir = getSubdirectory(file.type);
  const uniqueFilename = generateUniqueFilename(file.filename);
  const filePath = path.join(MEDIA_DIR, subdir, uniqueFilename);
  
  await fs.promises.writeFile(filePath, file.data);
  
  return {
    filename: uniqueFilename,
    originalName: file.filename,
    path: `/uploads/media/${subdir}/${uniqueFilename}`,
    size: file.data.length,
    type: file.type
  };
}

// Delete file from storage
export async function deleteFile(filePath: string) {
  const fullPath = path.join(process.cwd(), 'public', filePath);
  if (fs.existsSync(fullPath)) {
    await fs.promises.unlink(fullPath);
    return true;
  }
  return false;
}

// Get file info
export async function getFileInfo(filePath: string) {
  const fullPath = path.join(process.cwd(), 'public', filePath);
  if (fs.existsSync(fullPath)) {
    const stats = await fs.promises.stat(fullPath);
    return {
      size: stats.size,
      created: stats.birthtime,
      modified: stats.mtime
    };
  }
  return null;
} 