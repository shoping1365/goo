export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  const base = config.public.goApiBase
  const { getQuery } = await import('h3');
  const { proxy } = await import('../_utils/fetchProxy');
  const query = getQuery(event);
  const qs = Object.keys(query).length ? '?' + new URLSearchParams(query).toString() : '';
  return proxy(event, `${base}/api/post-categories${qs}`);
});