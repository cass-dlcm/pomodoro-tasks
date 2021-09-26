import styles from "./Mainpage.css";
//import notebook from './assets/notebook.svg';
//import clock from './assets/Clock.svg';
import { Clock } from "./Clock";
import { ToDoList }  from "./ToDoList";

//console.log(clock);
//console.log(notebook);

export const Mainpage = () => {
  return(
    <div>
      <section>
      <div class="countdown"><Clock></Clock></div>
        <section class = "timetext">
          <div class="settime">Set a Time</div>
          <div class="timeset">
            <div>Work Minutes</div>
            <div>Break Minutes</div>
          </div>
          <button class="timebutton">Begin!</button>
        </section>
      </section>

      <section>
        <div class="ToDoListApp"><ToDoList></ToDoList></div>
        <div class="notebook"><img src={'./assets/notebook.svg'} alt="notebook" /></div>
      </section>
    </div>
  );
}
//<div class="clock"><img src={clock} alt="clock" /></div>