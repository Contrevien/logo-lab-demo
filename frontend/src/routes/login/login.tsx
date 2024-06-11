import styles from './index.module.scss'
import Button from '@mui/material/Button'
import { useAuth0 } from '@auth0/auth0-react'
import { ReactComponent as LogoSvg } from '@/assets/icons/favicon.svg'

const Login = () => {
  const { loginWithRedirect } = useAuth0()

  return (
    <div className={styles.Login__Container}>
      <LogoSvg height={100} width={100} />
      <h1>Logo Lab Demo Chat</h1>
      <Button
        onClick={() => loginWithRedirect()}
        variant='contained'
        className={styles.Login__Button}
      >
        ログイン
      </Button>
    </div>
  )
}

export default Login
