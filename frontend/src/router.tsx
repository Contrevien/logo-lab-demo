import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom'
import { ROUTE_NAMES } from '@/utils/constants'
import Login from '@/routes/login/login'
import Chat from './routes/chat/chat'
import { useAuthContext } from './context/auth'

const Root = () => {
  const { user } = useAuthContext()

  return (
    <BrowserRouter>
      <Switch>
        <Route path={ROUTE_NAMES.LOGIN} component={Login} />
        {user && <Route path={ROUTE_NAMES.CHAT} component={Chat} />}
        <Redirect path='/' to={ROUTE_NAMES.LOGIN} />
      </Switch>
    </BrowserRouter>
  )
}

export default Root
