import Layout from './templates/Layout'
import Home from './templates/Home'
import { createGlobalStyle } from 'styled-components'


const GlobalStyle = createGlobalStyle`
  * {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
  }
`

const App = (): JSX.Element => (
  <div className="App">
    <GlobalStyle />
    <Layout>
      <Home />
    </Layout>
  </div>
)

export default App
