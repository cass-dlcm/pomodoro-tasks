import styles from "./Mainpage.css";
import { Clock } from "./Clock";
import {BreakClock} from "./Break Clock"
import { ToDoList }  from "./ToDoList";
import {gql, useQuery} from "@apollo/client";
import {useState} from "react";

export const Mainpage = () => {
    const [skip, setSkip] = useState(false);
    const [list, setList] = useState({id: -1});
    const [workTime, setWorkTime] = useState(true);
    const QUERYONE = gql`query GetListID {lists}`
    const {data, loading, error} = useQuery(QUERYONE, {skip: skip});
    if (data) {setList({id: data.lists[0]}); setSkip(true);}
    if (loading) return "Loading lists...";
    if (error) return `Error! ${error.message}`;

    function startWorkTime() {
        setWorkTime(true)
    }

    function startBreakTime() {
        setWorkTime(false)
    }

    return(
        <div>
            <section className={'timer-group'}>
                {workTime && <Clock />}
                {!workTime && <BreakClock/>}
                <div className="settime">Set a Time</div>
                <div className="timeset">
                    <div>25 Work Minutes</div>
                    <div>5 Break Minutes</div>
                </div>
                <div className="btn-section">
                    <button className="timebutton" onClick={startWorkTime}>
                        Begin!
                    </button>
                    <button className="timebutton" onClick={startBreakTime}>
                        Break!
                    </button>
                </div>
            </section>
            <section className={"todo-list-group"}>
                <ToDoList list={list} setList={setList}/>
                <div className="notebook"><img src={'./assets/notebook.svg'} height={"100%"} width={"100%"} alt="notebook" /></div>
            </section>
        </div>
    );
}