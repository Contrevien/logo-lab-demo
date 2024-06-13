import { Dispatch, SetStateAction, useEffect, useRef, useState } from 'react'
import cx from 'classnames'
import styles from './infiniteScroll.module.scss'

type InfiniteScrollProp = {
  page: number
  setPage: Dispatch<SetStateAction<number>>
  children: React.ReactNode
  shouldScroll: boolean | undefined
  className?: string
};

const InfiniteScroll = ({ page, setPage, children, shouldScroll, className }: InfiniteScrollProp) => {
  const scrollContainerRef = useRef<HTMLDivElement>(null)
  const [currentHeight, setCurrentHeight] = useState<number>(0)

  useEffect(() => {
    const handleScroll = () => {
      const scrollTop = scrollContainerRef.current?.scrollTop || 0
      if (scrollTop <= 20 && shouldScroll) {
        setPage((prevPage) => prevPage + 1)
      }
    }

    if (scrollContainerRef.current) {
      setCurrentHeight(scrollContainerRef.current.scrollHeight)
      scrollContainerRef.current.scrollTop = scrollContainerRef.current.scrollHeight
      scrollContainerRef.current.addEventListener('scroll', handleScroll)
    }
    return () => scrollContainerRef.current?.removeEventListener('scroll', handleScroll)
  }, [])

  useEffect(() => {
    if (scrollContainerRef.current) {
      setCurrentHeight(scrollContainerRef.current.scrollHeight - currentHeight)
      scrollContainerRef.current.scrollTop = scrollContainerRef.current.scrollHeight - currentHeight
    }
  }, [page])

  return <div className={cx(styles.InfiniteScroll__Container, className)} ref={scrollContainerRef}>{children}</div>
}

export default InfiniteScroll
