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
import {gql, useQuery} from "@apollo/client";
import {useState} from "react";



export const Mainpage = (props) => {
    const [skip, setSkip] = useState(false);
    const [list, setList] = useState({id: -1});
    const QUERYONE = gql`query GetListID {lists}`
    const {data, loading, error} = useQuery(QUERYONE, {skip: skip});
    if (data) {setList({id: data.lists[0]}); setSkip(true);}
    if (loading) return "Loading lists...";
    if (error) return `Error! ${error.message}`;

    return(
        <div>
            <div>
                <section>
                    <div className="countdown">
                        <Route exact path="/clock" component={Clock} />
                        <Route exact path="/breakclock" component={BreakClock} />
                        <Clock />
                    </div>
                    <section className = "timetext">
                        <div className="settime">Set a Time</div>
                        <div className="timeset">
                            <div>25 Work Minutes</div>
                            <div>5 Break Minutes</div>
                        </div>
                        <div className="btn-section">
                            <button className="timebutton">
                                <a href="">
                                    <Link to="/clock"> Begin! </Link>
                                </a>
                            </button>
                            <button className="timebutton">
                                <a href="">
                                    <Link to="/breakclock"> Break! </Link>
                                </a>
                            </button>
                        </div>
                    </section>
                </section>
                <section>
                    <div className="ToDoListApp"><ToDoList list={list} setList={setList}/></div>
                    <div className="notebook"><img src={'./assets/notebook.svg'} alt="notebook" /></div>
                </section>
            </div>
        </div>
    );

}