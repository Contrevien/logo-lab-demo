import { createApi } from '@reduxjs/toolkit/query/react'

import { fetcherQuery } from './fetcher'

export const API_TAGS = {
  MESSAGE: 'MESSAGE'
} as const

export const api = createApi({
  reducerPath: 'api',
  tagTypes: [...Object.values(API_TAGS)],
  baseQuery: (args, api, extraOptions = {}) =>
    fetcherQuery(args, api, extraOptions),
  endpoints: () => ({}),
})
