import { defineEventHandler } from 'h3'

export default defineEventHandler(async (_event) => {
  // احراز هویت غیرفعال شده است - وضعیت سیستم mock برمی‌گرداند
  return {
    status: 'healthy',
    timestamp: new Date().toISOString(),
    services: {
      image_compression: {
        status: 'running',
        active_jobs: 2,
        completed_jobs: 15,
        failed_jobs: 1,
        queue_size: 5
      },
      video_compression: {
        status: 'running',
        active_jobs: 1,
        completed_jobs: 8,
        failed_jobs: 0,
        queue_size: 3
      },
      storage: {
        status: 'healthy',
        total_space: 107374182400, // 100GB
        used_space: 21474836480,   // 20GB
        available_space: 85899345920 // 80GB
      },
      database: {
        status: 'connected',
        connection_pool: {
          active: 5,
          idle: 10,
          max: 20
        }
      }
    },
    performance: {
      cpu_usage: 25.5,
      memory_usage: 45.2,
      disk_io: {
        read_mbps: 12.5,
        write_mbps: 8.3
      }
    }
  }
}) 