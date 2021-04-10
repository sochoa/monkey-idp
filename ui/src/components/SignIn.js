import React, { useState } from 'react'
import { Alert, Form, Col, Button, Modal } from 'react-bootstrap'

export default function SignIn () {
  const [showModal, setShowModal] = useState(false)
  const [authResult, setAuthResult] = useState('')
  const [authStatus, setAuthStatus] = useState('success')

  function handleSubmit (event) {
    event.preventDefault()
    fetch('/api/v1/auth', {
      method: 'POST',
      body: JSON.stringify({
        password: document.querySelector('#password').value,
        userid: document.querySelector('#userid').value
      }),
      headers: { 'Content-Type': 'application/json' }
    })
      .then(function (response) {
        if (response.ok) {
          setAuthResult('Authorized!')
          setAuthStatus('light')
        } else {
          setAuthResult('Unauthorized')
          setAuthStatus('danger')
        }
        setShowModal(true)
      })
  }

  return (
    <div>
      <div className="row">
        <div className="col-md-12">
          <Form onSubmit={handleSubmit}>
            <Form.Row>
              <Form.Group as={Col} controlId="userid" column lg="5">
                <Form.Label>User ID</Form.Label>
                <Form.Control required autoFocus type="text" placeholder="Enter User ID" />
              </Form.Group>
            </Form.Row>

            <Form.Row>
              <Form.Group as={Col} controlId="password" column lg="5">
                <Form.Label>Password</Form.Label>
                <Form.Control required type="password" placeholder="Password" />
              </Form.Group>
            </Form.Row>

            <Button variant="primary" type="submit">
              Sign In
            </Button>
          </Form>
        </div>
      </div>
      <Modal
        show={showModal}
        size="lg"
        aria-labelledby="submit-modal-title"
        centered
      >
        <Modal.Body>
          <p><Alert variant={authStatus}>{authResult}</Alert></p>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="primary" onClick={() => setShowModal(false)}>OK</Button>
        </Modal.Footer>
      </Modal>
    </div>
  )
}
