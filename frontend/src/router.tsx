import { BrowserRouter, Route, Switch } from 'react-router-dom'
import { ROUTE_NAMES } from '@/utils/constants'
import Auth from '@/routes/auth'

const Root = () => {
  return (
    <BrowserRouter>
      <Switch>
        <Route path={ROUTE_NAMES.AUTH} component={Auth} />
      </Switch>
    </BrowserRouter>
  )
}

export default Root
