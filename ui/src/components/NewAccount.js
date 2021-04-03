import React from 'react'
import { Form, Col, Button } from 'react-bootstrap'

const NewAccount = () => {
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
              <Form.Group as={Col} controlId="formEmail" column lg="5">
                <Form.Label>Email</Form.Label>
                <Form.Control type="email" placeholder="Enter Email" />
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

export default NewAccount
