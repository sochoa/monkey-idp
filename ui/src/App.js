import Header from './components/Header.js'
import RequestRouter from './components/RequestRouter.js'

import Home from './components/Home.js'
import NewAccount from './components/NewAccount.js'
import SignIn from './components/SignIn.js'

function App() {
  const availableRoutes = [
    {
      component: Home,
      key: "home",
      path: "/",
      label: "Home"
    },
    {
      component: NewAccount,
      key: "new-account",
      path: "/new-account",
      label: "Create an Account"
    },
    {
      component: SignIn,
      key: "sign-in",
      path: "/sign-in",
      label: "Sign In"
    },
  ]
  return (
    <div className='container'>
      <RequestRouter routes={availableRoutes} />
      <Header routes={availableRoutes} />
    </div>
  );
}

export default App;
