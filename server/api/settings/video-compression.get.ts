// احراز هویت غیرفعال شده است - تنظیمات فشرده‌سازی ویدئو
export default defineEventHandler(async (event) => {
  return [
    { key: 'video_compression.quality', value: '85', category: 'video_compression', type: 'string' },
    { key: 'video_compression.format', value: 'mp4', category: 'video_compression', type: 'string' },
    { key: 'video_compression.codec', value: 'h264', category: 'video_compression', type: 'string' },
    { key: 'video_compression.two_pass', value: 'false', category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.remove_metadata', value: 'true', category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.create_thumbnail', value: 'true', category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.frame_rate', value: '30', category: 'video_compression', type: 'string' },
    { key: 'video_compression.custom_width', value: '1920', category: 'video_compression', type: 'string' },
    { key: 'video_compression.custom_height', value: '1080', category: 'video_compression', type: 'string' },
    { key: 'video_compression.custom_bitrate', value: '2000', category: 'video_compression', type: 'string' },
    { key: 'video_compression.worker_count', value: '4', category: 'video_compression', type: 'string' },
    { key: 'video_compression.frame_analysis_mode', value: 'auto', category: 'video_compression', type: 'string' },
    { key: 'video_compression.enabled', value: 'true', category: 'video_compression', type: 'boolean' },
    { key: 'video_compression.start_hour', value: '0', category: 'video_compression', type: 'string' },
    { key: 'video_compression.end_hour', value: '6', category: 'video_compression', type: 'string' }
  ]
})