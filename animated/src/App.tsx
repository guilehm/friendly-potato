import { createGlobalStyle } from 'styled-components'
import Layout from './templates/Layout'
import Home from './templates/Home'
import Login from './templates/Login'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from 'react-router-dom'
import { ChakraProvider } from '@chakra-ui/react'


const GlobalStyle = createGlobalStyle`
  * {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
    font-family: 'Raleway', sans-serif;
  }
`

const App = (): JSX.Element => (
  <div className="App">
    <GlobalStyle />
    <ChakraProvider>
      <Router>
        <Switch>

          <Route path="/login">
            <Login />
          </Route>

          <Route path="/">
            <Layout>
              <Home />
            </Layout>
          </Route>

        </Switch>
      </Router>
    </ChakraProvider>
  </div>
)

export default App
