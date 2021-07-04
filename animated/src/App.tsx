import Layout from './templates/Layout'
import Home from './templates/Home'


const App = (): JSX.Element => (
  <div className="App">
    <Layout>
      <Home />
    </Layout>
  </div>
)

export default App
