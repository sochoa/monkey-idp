import PropTypes from 'prop-types'

import {
  BrowserRouter as Router,
  Switch,
  Route,
  useParams,
} from "react-router-dom";

import {
  Navbar,
  Nav,
  NavDropdown,
  Form,
  FormControl,
  Button
} from 'react-bootstrap'

import Home from './Home.js'
import NewAccount from './NewAccount.js'
import SignIn from './SignIn.js'

const Header = ({title}) => {
  return (
    <div>
      <div className="row">
        <div className="col-md-12">
          <Router>
            <Navbar bg="dark" variant="dark" expand="lg" sticky="top">
              <Navbar.Brand>{title}</Navbar.Brand>
              <Navbar.Toggle aria-controls="navbar-main" />
              <Navbar.Collapse id="navbar-main">
                <Nav className="mr-auto">
                  <Nav.Link href="/">Home</Nav.Link>
                  <Nav.Link href="/new-account">Create an Account</Nav.Link>
                  <Nav.Link href="/sign-in">Sign In</Nav.Link>
                </Nav>
              </Navbar.Collapse>
            </Navbar>
            <br />
            <Switch>
              <Route exact path="/">
                <Home />
              </Route>
              <Route path="/new-account">
                <NewAccount />
              </Route>
              <Route path="/sign-in">
                <SignIn />
              </Route>
            </Switch>
          </Router>
        </div>
      </div>
    </div>
  )
}

Header.defaultProps = {
  title: 'Monkey IDP',
}

Header.propTypes = {
  title: PropTypes.string
}

export default Header
