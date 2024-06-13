/* eslint-disable no-console */
import {
  type BaseQueryFn,
  type FetchArgs,
  fetchBaseQuery,
} from '@reduxjs/toolkit/query/react'

const BASE_QUERY = (baseUrl = import.meta.env.VITE_API_URL) =>
  fetchBaseQuery({
    baseUrl,
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('accessToken')
      if (token) {
        headers.set('Authorization', token)
      }

      return headers
    },
  })

type BaseQueryError = {
  message: string;
};

type BaseQueryFnArgs = Parameters<
  BaseQueryFn<string | FetchArgs, unknown, BaseQueryError, { baseUrl?: string }>
>;

export const fetcherQuery = async (...args: BaseQueryFnArgs) => {
  const [baseArgs, api, extraOptions] = args
  const { baseUrl } = extraOptions
  const queryBuilder = BASE_QUERY(baseUrl)
  let result: Awaited<ReturnType<typeof queryBuilder>>

  try {
    result = await queryBuilder(baseArgs, api, extraOptions)
  } catch (err) {
    return {
      error: {
        message: 'Unexpected error occurred',
      },
    }
  }

  const { error } = result
  if (error && 'data' in error) {
    // this should not happen, this means server only send null back to us with error status code
    if (!error.data) {
      return {
        error: {
          message: 'Unexpected error occurred',
        },
      }
    }

    const _error = error.data as BaseQueryError

    return {
      error: _error,
    }
  }

  return result
}
