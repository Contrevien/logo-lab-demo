import { makeQueryString } from '@/utils/utils'
import { api as coreApi } from './api'

type PostMessageBody = {
  text: string;
  sentByName: string;
  sentById: string;
};

export type Message = {
  _id: string;
  text: string;
  sentByName: string;
  sentById: string;
  createdAt: number;
};

type GetMessageResponse = {
  messages: Message[];
  hasMore: boolean;
};

type GetMessageArgs = {
  page: number;
};

const api = coreApi.injectEndpoints({
  endpoints: (builder) => ({
    postMessage: builder.mutation<unknown, PostMessageBody>({
      query: (payload) => ({
        url: '/messages',
        method: 'POST',
        body: payload,
      }),
    }),
    getMessages: builder.query<GetMessageResponse, GetMessageArgs>({
      query: ({ page }) => ({
        url: `/messages?${makeQueryString({ page })}`,
        method: 'GET',
      }),
    }),
  }),
})

export const { useGetMessagesQuery, usePostMessageMutation } = api
