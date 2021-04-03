import React from 'react'

import Header from './components/Header.js'
import RequestRouter from './components/RequestRouter.js'
import AVAILABLE_ROUTES from './definitions/routes.js'

function App () {
  return (
    <div className='container'>
      <RequestRouter routes={AVAILABLE_ROUTES} />
      <Header routes={AVAILABLE_ROUTES} />
    </div>
  )
}

export default App
