import Header from './components/Header.js'
import RequestRouter from './components/RequestRouter.js'

function App() {
  return (
    <div className='container'>
      <RequestRouter />
      <Header routes="{RequestRouter.AvailableRoutes}"/>
    </div>
  );
}

export default App;
