import PropTypes from 'prop-types'

import { Navbar, Nav } from 'react-bootstrap'

const Header = ({ routes, title }) => {
  return (
    <div>
      <div className="row">
        <div className="col-md-12">
          <Navbar bg="dark" variant="dark" expand="md" sticky="top">
            <Navbar.Brand href="/">{title}</Navbar.Brand>
            <Nav className="ml-auto">
              {routes.map((route) => (
                <Nav.Link href={`${route.path}`}>{route.label}</Nav.Link>
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
  routes: []
}

Header.propTypes = {
  title: PropTypes.string,
  routes: PropTypes.array
}

export default Header
