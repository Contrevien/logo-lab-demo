import { ReactComponent as LogoSvg } from '@/assets/icons/favicon.svg'
import styles from './loading.module.scss'

const Loading = () => {
  return <div className={styles.Loading}><LogoSvg height={100} width={100} /></div>
}

export default Loading
