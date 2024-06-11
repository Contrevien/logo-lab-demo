import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom'
import { ROUTE_NAMES } from '@/utils/constants'
import Login from '@/routes/login/login'

const Root = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route path={ROUTE_NAMES.LOGIN} component={Login} />
        <Redirect path='/' to={ROUTE_NAMES.LOGIN} />
      </Switch>
    </BrowserRouter>
  )
}

export default Root
