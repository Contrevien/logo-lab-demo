import { useEffect, useState } from 'react'
import styles from './chat.module.scss'
import InfiniteScroll from './components/infiniteScroll'
import Message from './components/message'
import Input from './components/input'
import { type Message as MessageType, useGetMessagesQuery } from '@/services/api/chat'
import { useAuth0 } from '@auth0/auth0-react'

const Chat = () => {
  const [page, setPage] = useState(1)
  const [messages, setMessages] = useState<MessageType[]>([])
  const { user } = useAuth0()

  const { data, isLoading } = useGetMessagesQuery({
    page,
  })

  useEffect(() => {
    if (data?.messages) {
      setMessages([...messages, ...data.messages])
    }
  }, [data])

  return (
    <>
      <InfiniteScroll page={page} setPage={setPage} shouldScroll={data && data?.hasMore}>
        <div className={styles.Chat__MessagesWrapper}>
          {messages.map(m => <Message
            key={m._id}
            sentBy={m.sentByName}
            text={m.text}
            self={user?.sub === m.sentByID}
            createdAt={m.createdAt} />
          )}
        </div>
      </InfiniteScroll>
      <Input />
    </>
  )
}

export default Chat
