import { defineEventHandler, getQuery } from 'h3';
import { proxy } from '../_utils/fetchProxy';

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const query = getQuery(event)
  // پشتیبانی از پارامتر all=1 (قبلاً اشتباهاً 'all1' لاگ می‌شد)
  if (query.all === '1' || query.all === 1 || query.all === true || query.all === 'true') {
    query.all = '1'
  }
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query as Record<string, string>).toString() : '';
  return proxy(event, `${base}/api/product-categories${qs}`);
});
