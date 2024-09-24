import type { LayoutServerLoad } from './$types'
import { getSession } from '$lib/session'

export const load: LayoutServerLoad = async ({ locals }) => {
  const { user } = locals
  const session = getSession(user)

  // load could also return additional data
  // other than the session, such as site config
  return { session }
}