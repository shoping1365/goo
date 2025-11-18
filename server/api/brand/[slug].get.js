import { defineEventHandler } from 'h3'

export default defineEventHandler(async (event) => {
  const { slug } = event.context.params;
  const base = getGoApiBaseUrl();

  try {
    const brands = await $fetch(`${base}/api/brands`);
    const brand = (Array.isArray(brands) ? brands : brands.data || []).find(b => b.slug === slug);
    if (!brand) {
      throw createError({ statusCode: 404, statusMessage: 'Brand not found' });
    }
    return brand;
  } catch (err) {
    // Error fetching brand by slug
    throw createError({ statusCode: 500, statusMessage: 'Internal server error' });
  }
});