import { createGlobalStyle } from "styled-components"
import Layout from "./templates/Layout"
import Games from "./templates/Games"
import Lumber from "./templates/Lumber"
import RPG from "./components/RPG"
import Login from "./templates/Login"
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"
import { ChakraProvider } from "@chakra-ui/react"
import GameDetail from "./templates/GameDetail"
import RogueLike from "./components/RogueLike"


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

          <Route path="/login/">
            <Layout>
              <Login />
            </Layout>
          </Route>

          <Route path="/games/:slug/">
            <Layout>
              <GameDetail />
            </Layout>
          </Route>

          <Route path="/lumber/" render={(props) =>
            <Layout>
              <Lumber {...props} />
            </Layout>}>
          </Route>

          <Route path="/games/">
            <Layout>
              <Games />
            </Layout>
          </Route>

          <Route path="/rpg/">
            <Layout>
              <RPG />
            </Layout>
          </Route>

          <Route path="/">
            <RogueLike />
          </Route>

        </Switch>
      </Router>
    </ChakraProvider>
  </div>
)

export default App
