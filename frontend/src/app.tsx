import Root from './router'
import { Auth0Provider } from '@auth0/auth0-react'
import { ROUTE_NAMES } from './utils/constants'
import { ThemeProvider } from '@mui/material'
import theme from './utils/theme'
import { AuthProvider } from './context/auth'

const AUTH0_PROPS = {
  domain: import.meta.env.VITE_OAUTH_DOMAIN,
  clientId: import.meta.env.VITE_OAUTH_CLIENT_ID,
  authorizationParams: {
    redirect_uri: `https://${window.location.hostname}${ROUTE_NAMES.LOGIN}`,
    ui_locale: 'ja',
  },
}

const App = () => {
  return (
    <Auth0Provider {...AUTH0_PROPS}>
      <ThemeProvider theme={theme}>
        <AuthProvider>
          <Root />
        </AuthProvider>
      </ThemeProvider>
    </Auth0Provider>
  )
}

export default App
