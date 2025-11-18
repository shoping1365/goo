import { defineEventHandler, getRouterParam } from 'h3';
import { proxy } from "../../../_utils/fetchProxy";

export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, 'id')
  if (!id) {
    throw new Error('ID is required')
  }
  return proxy(event, `/api/admin/fraud/cases/${id}`);
});

