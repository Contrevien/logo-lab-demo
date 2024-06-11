import { useAuth0 } from '@auth0/auth0-react'
import {
  Dispatch,
  ReactNode,
  SetStateAction,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react'

type AuthContextProps = {
  user: any;
  setUser: Dispatch<SetStateAction<any>>;
};

const AuthContext = createContext<AuthContextProps>({
  user: null,
  setUser: () => {},
})

type AuthProviderProps = {
  children: ReactNode;
};

export function AuthProvider (props: AuthProviderProps) {
  const { children } = props
  const [user, setUser] = useState<any>()
  const {
    user: auth0User,
    isAuthenticated,
    getAccessTokenWithPopup,
  } = useAuth0()

  useEffect(() => {
    const getUserMetadata = async () => {
      try {
        const accessToken = await getAccessTokenWithPopup({
          authorizationParams: {
            audience: `https://${import.meta.env.VITE_OAUTH_DOMAIN}/api/v2/`,
            scope: 'read:current_user',
            prompt: 'consent',
          },
        })

        const userDetailsByIdUrl = `https://${
          import.meta.env.VITE_OAUTH_DOMAIN
        }/api/v2/users/${auth0User?.sub}`

        const metadataResponse = await fetch(userDetailsByIdUrl, {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        })

        const { user_metadata } = await metadataResponse.json()

        setUser(user_metadata)
      } catch (e) {
        console.log(e)
      }
    }

    if (isAuthenticated && auth0User) getUserMetadata()
  }, [getAccessTokenWithPopup, auth0User?.sub, isAuthenticated])

  return (
    <AuthContext.Provider
      // eslint-disable-next-line react/jsx-no-constructed-context-values
      value={{
        user,
        setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export const useAuthContext = () => useContext(AuthContext)
