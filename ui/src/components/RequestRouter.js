import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom'

const RequestRouter = ({ routes }) => {
  return (
    <Router>
      <Switch>
        {routes.map(route => <Route exact key={route.key} path={`${route.path}`}></Route>)}
      </Switch>
    </Router>
  )
}

export default RequestRouter
