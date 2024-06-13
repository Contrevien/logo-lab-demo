import { createApi } from '@reduxjs/toolkit/query/react'

import { fetcherQuery } from './fetcher'

export const api = createApi({
  reducerPath: 'api',
  baseQuery: (args, api, extraOptions = {}) =>
    fetcherQuery(args, api, extraOptions),
  endpoints: () => ({}),
})
