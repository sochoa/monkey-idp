import {
  BrowserRouter as Router,
  Switch,
  Route,
  useParams,
} from "react-router-dom";

import Home from './Home.js'
import NewAccount from './NewAccount.js'
import SignIn from './SignIn.js'


const RequestRouter = ({routes}) => {
  const AvailableRoutes = [
    {
      component: Home,
      path: "/",
      label: "Home"
    },
    {
      component: NewAccount,
      path: "/new-account",
      label: "Create an Account"
    },
    {
      component: SignIn,
      path: "/sign-in",
      label: "Sign In"
    },
  ]
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <Home />
        </Route>
        {routes.map(funcion(route, i) {
          const RouteComponent = AvailableRoutes[route].component
          const RoutePath = AvailableRoutes[route].path
          return <Route exact path="{RoutePath}"><RouteComponent /></Route>
        })}
      </Switch>
    </Router>
  )
}

export default RequestRouter
