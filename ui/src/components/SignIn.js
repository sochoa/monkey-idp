import React from 'react'
import { Form, Col, Button } from 'react-bootstrap'

export default function SignIn () {
  const fields = {
    userid: null,
    password: null,
    email: null
  }

  async function handleSubmit (event) {
    event.preventDefault()
    fetch('/api/v1/user', {
      method: 'POST',
      body: JSON.stringify({}),
      headers: { 'Content-Type': 'application/json' }
    })
      .then(res => res.json())
    // .then(json => setUser(json.user))
  }

  return (
    <div>
      <div className="row">
        <div className="col-md-12">
          <Form onSubmit={handleSubmit}>
            <Form.Row>
              <Form.Group as={Col} controlId="userid" column lg="5">
                <Form.Label>User ID</Form.Label>
                <Form.Control autoFocus type="text" placeholder="Enter User ID" value={fields.userid} />
              </Form.Group>
            </Form.Row>

            <Form.Row>
              <Form.Group as={Col} controlId="password" column lg="5">
                <Form.Label>Password</Form.Label>
                <Form.Control type="password" placeholder="Password" />
              </Form.Group>
            </Form.Row>

            <Button variant="primary" type="submit">
              Sign In
            </Button>
          </Form>
        </div>
      </div>
    </div>
  )
}
