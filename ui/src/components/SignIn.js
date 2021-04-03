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
              <Form.Group as={Row} controlId="signInFormUserId">
                <Form.Label column sm="2" style={{ textAlign: 'right' }}>User ID</Form.Label>
                <Col sm="5">
                  <Form.Control type="text" placeholder="Your Email" />
                </Col>
              </Form.Group>
              <Form.Group as={Row} controlId="signInFormPassword">
                <Form.Label column sm="2" style={{ textAlign: 'right' }}>Password</Form.Label>
                <Col sm="5">
                  <Form.Control type="password" placeholder="Your Password" />
                </Col>
              </Form.Group>
              <Form.Row>
                <Col sm={{ span: 5, offset: 2 }}>
                  <Button variant="primary" type="submit">
                    Submit
                  </Button>
                </Col>
              </Form.Row>
            </Form>
          </div>
        </div>
      </div>
    )
  }
}

export default SignIn
