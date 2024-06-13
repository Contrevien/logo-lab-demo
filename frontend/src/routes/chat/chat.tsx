import styles from './index.module.scss'
import InfiniteScroll from './components/infiniteScroll/infiniteScroll'
import { useState } from 'react'

const Chat = () => {
  const [page, setPage] = useState(0)

  const Block = () => <div className={styles.Block} />

  let blocks = [
    <Block />,
    <Block />,
    <Block />,
    <Block />,
    <Block />,
    <Block />,
    <Block />,
    <Block />,
  ]

  for (let i = 0; i < page; i++) blocks = blocks.concat(blocks)

  return (
    <div className={styles.Chat__Container}>
      <InfiniteScroll page={page} setPage={setPage} data={blocks} />
    </div>
  )
}

export default Chat
