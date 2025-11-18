import { eventHandler, setResponseHeaders, getMethod } from 'h3'

export default eventHandler((event) => {
     // موقتاً غیرفعال برای رفع مشکلات
     return;

     // کدهای زیر فعلاً غیرفعال هستند
     /*
     setResponseHeaders(event, {
          'Access-Control-Allow-Origin': '*',
          'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
          'Access-Control-Allow-Headers': 'Content-Type, Authorization',
          'Access-Control-Allow-Credentials': 'true'
     })

     if (getMethod(event) === 'OPTIONS') {
          event.node.res.statusCode = 200
          event.node.res.end()
     }
     */
})
