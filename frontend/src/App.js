import { Home } from "./components/Home";
import { SignUp } from "./components/SignUp";
import { Mainpage } from "./components/Mainpage";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

function App() {
  return (
    <div >
      <Router>
        <Route exact path="/" component={Home} />
        <Route exact path="/signup" component={SignUp} />
        <Route exact path="/mainpage" component={Mainpage} />
        </Router>

    </div>
  );
}

export default App;