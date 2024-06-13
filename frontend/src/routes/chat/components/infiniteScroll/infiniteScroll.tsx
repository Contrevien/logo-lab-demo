import { Dispatch, SetStateAction, useEffect } from 'react'
import styles from './index.module.scss'

type InfiniteScrollProp = {
  page: number;
  setPage: Dispatch<SetStateAction<number>>;
  data: React.ReactNode[];
};

const InfiniteScroll = ({ page, setPage, data }: InfiniteScrollProp) => {
  useEffect(() => {
    window.addEventListener('scroll', handleScroll)
    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  const handleScroll = () => {
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
      setPage(page + 1)
    }
  }

  return <div className={styles.InfiniteScroll__Container}>{data}</div>
}

export default InfiniteScroll
