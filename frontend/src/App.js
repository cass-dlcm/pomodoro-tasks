import { Home } from "./components/Home";
import { SignUp } from "./components/SignUp";
import { Clock } from "./components/Clock"

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
        <Route exact path="/clock" component={Clock} />
        </Router>
      
    </div>
  );
}

export default App;
