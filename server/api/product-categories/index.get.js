export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { getQuery } = await import('h3');
  const { proxy } = await import('../_utils/fetchProxy');
  const query = getQuery(event);
  // پشتیبانی از پارامتر all=1 (قبلاً اشتباهاً 'all1' لاگ می‌شد)
  if (query.all === '1' || query.all === 1 || query.all === true || query.all === 'true') {
    query.all = '1'
  }
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query).toString() : '';
  return proxy(event, `${base}/api/product-categories${qs}`);
});