import Home from '../components/Home.js'
import NewAccount from '../components/NewAccount.js'
import SignIn from '../components/SignIn.js'

export const AVAILABLE_ROUTES = [
  {
    component: Home,
    key: 'home',
    path: '/',
    label: 'Home'
  },
  {
    component: NewAccount,
    key: 'new-account',
    path: '/new-account',
    label: 'Create an Account'
  },
  {
    component: SignIn,
    key: 'sign-in',
    path: '/sign-in',
    label: 'Sign In'
  }
]
