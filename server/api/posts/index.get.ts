import { defineEventHandler, getQuery, parseCookies } from 'h3';

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  
  // Get all cookies from the request
  const cookies = parseCookies(event);
  const cookieHeader = Object.entries(cookies)
    .map(([key, value]) => `${key}=${value}`)
    .join('; ');

  const query = getQuery(event);
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : '';
  
  try {
    const response = await $fetch(`${base}/api/posts${qs}`, {
      headers: {
        'Cookie': cookieHeader
      }
    });
    
    return response;
  } catch (error) {
    console.error('Posts index API - Error:', error);
    throw error;
  }
});
