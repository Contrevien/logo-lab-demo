import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom'
import { ROUTE_NAMES } from '@/utils/constants'
import Login from '@/routes/login/login'
import Callback from './routes/callback/callback'

const Root = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route path={ROUTE_NAMES.LOGIN} component={Login} />
        <Route path={ROUTE_NAMES.CALLBACK} component={Callback} />
        <Redirect path='/' to={ROUTE_NAMES.LOGIN} />
      </Switch>
    </BrowserRouter>
  )
}

export default Root
