import React from 'react'
import PropTypes from 'prop-types'

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

RequestRouter.defaultProps = {
  routes: []
}

RequestRouter.propTypes = {
  routes: PropTypes.array
}

export default RequestRouter
