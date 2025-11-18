import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  // احراز هویت غیرفعال شده است - لیست ویدیوها mock برمی‌گرداند
  const sampleVideos = [
    {
      id: 1,
      file_name: 'sample_video_1.mp4',
      url: '/uploads/videos/sample_video_1.mp4',
      thumbnail: '/uploads/thumbnails/sample_video_1.jpg',
      size: 52428800, // 50MB
      duration_seconds: 120,
      width: 1920,
      height: 1080,
      format: 'mp4',
      bitrate_kbps: 5000
    },
    {
      id: 2,
      file_name: 'sample_video_2.webm',
      url: '/uploads/videos/sample_video_2.webm',
      thumbnail: '/uploads/thumbnails/sample_video_2.jpg',
      size: 31457280, // 30MB
      duration_seconds: 90,
      width: 1280,
      height: 720,
      format: 'webm',
      bitrate_kbps: 3000
    }
  ]
  
  return {
    success: true,
    data: sampleVideos
  }
}) 