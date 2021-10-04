import { Home } from "./components/Home";
import { SignUp } from "./components/SignUp";
import { Mainpage } from "./components/Mainpage";
import { BrowserRouter as Router, Route } from "react-router-dom";
import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";

function App() {
    const clientA = new ApolloClient({
        uri: 'http://localhost:8080/query',
        cache: new InMemoryCache()
    });

    let clientB = clientA;

    const setJWT = (val) => {
        clientB = new ApolloClient({
            uri: 'http://localhost:8080/query',
            cache: new InMemoryCache(),
            headers: {
                authorization: "Bearer " + val
            }
        });
    }
    return (
        <Router forceRefresh={true}>
            <ApolloProvider client={clientA}>
                <Route exact path="/" render={() => <Home sendJWT={setJWT} />} />
                <Route exact path="/signup" component={SignUp} />
            </ApolloProvider>
            <ApolloProvider client={clientB}>
                <Route exact path="/mainpage" component={Mainpage} />
            </ApolloProvider>
        </Router>
    );
}

export default App;