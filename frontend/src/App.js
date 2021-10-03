import { Home } from "./components/Home";
import { SignUp } from "./components/SignUp";
import { Mainpage } from "./components/Mainpage";
import { BrowserRouter as Router, Route } from "react-router-dom";

function App(props) {
    const getJWT = (val) => {
        props.setJWT(val);
    }
    return (
        <Router forceRefresh={true}>
            <Route exact path="/" render={() => <Home sendJWT={getJWT} />} />
            <Route exact path="/signup" component={SignUp} />
            <Route exact path="/mainpage" component={Mainpage} />
        </Router>
    );
}

export default App;