import React from 'react'
import PropTypes from 'prop-types'

import {
  BrowserRouter as Router,
  Switch,
  Route
} from 'react-router-dom'

import Home from './Home.js'
import NewAccount from './NewAccount.js'
import SignIn from './SignIn.js'

const RequestRouter = ({ routes }) => {
  return (
    <Router>
      <Switch>
        <Route exact path="/">
          <Home />
        </Route>
        <Route exact path="/sign-in">
          <SignIn />
        </Route>
        <Route exact path="/new-account">
          <NewAccount />
        </Route>
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
