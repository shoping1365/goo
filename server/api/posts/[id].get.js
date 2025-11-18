import { defineEventHandler, parseCookies } from 'h3'

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  
  // Get all cookies from the request
  const cookies = parseCookies(event);
  const cookieHeader = Object.entries(cookies)
    .map(([key, value]) => `${key}=${value}`)
    .join('; ');

  try {
    // همیشه با کوکی درخواست بزن (حتی اگر خالی باشد)
    const response = await $fetch(`${base}/api/posts/${id}${event.req.url.includes('preview=1') ? '?preview=1' : ''}`, {
      headers: {
        'Cookie': cookieHeader || 'mysession='
      }
    });
    
    return response;
  } catch (error) {
    console.error('Posts API - Error:', error);
    throw error;
  }
})