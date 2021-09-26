import styles from "./Mainpage.css";
//import notebook from './assets/notebook.svg';
//import clock from './assets/Clock.svg';
import { Clock } from "./Clock";
import {BreakClock} from "./Break Clock"
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import { ToDoList }  from "./ToDoList";

export const Mainpage = () => {
  
  return(
    <Router>
      <div>
      <section>
      <div class="countdown">
        <Route exact path="/clock" component={Clock} />
        <Route exact path="/breakclock" component={BreakClock} />
        <Clock />
      </div>
        <section class = "timetext">
          <div class="settime">Set a Time</div>
          <div class="timeset">
            
            <div>25 Work Minutes</div>
            <div>5 Break Minutes</div>
          </div>
          <div class="btn-section">
          <button class="timebutton">
            <a href="">
           <Link to="/clock"> Begin! </Link>
            </a>
          </button>
          <button class="timebutton">
            <a href="">
            <Link to="/breakclock"> Break! </Link>
            </a>
            </button>
          </div>
        </section>
      </section>

      <section>
        <div class="ToDoListApp"><ToDoList></ToDoList></div>
        <div class="notebook"><img src={'./assets/notebook.svg'} alt="notebook" /></div>
      </section>
    </div>
    </Router>
  );
  
}