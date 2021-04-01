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

const Header = ({props}) => {
  return (
    <div>
      <div className="row">
        <div className="col-md-12">
          <Navbar bg="dark" variant="dark" expand="md" sticky="top">
            <Navbar.Brand href="/">{props.title}</Navbar.Brand>
            <Nav className="ml-auto">
              {props.routes.map((route) => (
                <Nav.Link href="{route.path}">{route.label}</Nav.Link>
              ))}
            </Nav>
          </Navbar>
          <br />
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
