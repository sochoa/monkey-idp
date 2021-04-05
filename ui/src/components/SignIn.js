import React from 'react'
import { Form, Row, Col, Button } from 'react-bootstrap'

class SignIn extends React.Component {
  constructor (props) {
    super(props)
    this.state = { value: '' }
    this.handleChange = this.handleChange.bind(this)
    this.handleSubmit = this.handleSubmit.bind(this)
  }

  handleChange (event) {
    this.setState({ value: event.target.value })
  }

  handleSubmit (event) {
    // TODO:  POST to API
    alert('A name was submitted: ' + this.state.value)
    event.preventDefault()
  }

  render () {
    return (
      <div>
        <div className="row">
          <div className="col-md-12">
            <Form>
              <Form.Row>
                <Form.Group as={Col} controlId="formUserId" column lg="5">
                  <Form.Label>User ID</Form.Label>
                  <Form.Control type="text" placeholder="Enter User ID" />
                </Form.Group>
              </Form.Row>

              <Form.Row>
                <Form.Group as={Col} controlId="formPassword" column lg="5">
                  <Form.Label>Password</Form.Label>
                  <Form.Control type="password" placeholder="Password" />
                </Form.Group>
              </Form.Row>

              <Button variant="primary" type="submit">
                Submit
              </Button>
            </Form>
          </div>
        </div>
      </div>
    )
  }
}

export default SignIn
