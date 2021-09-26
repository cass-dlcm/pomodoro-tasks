import styles from "./Home.module.scss";
import { useHistory } from "react-router";

import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom";

export const Home = () => {
    let history = useHistory();

    return (
        <div>
            <div className={styles.background}>
                <div className={styles.sec1}>
                <img src="./assets/Laptop.png" alt="" />
                </div>
                <div className={styles.sec2}>
                <img src="./assets/StickyNote.png" alt="" />
                <div className={styles.loginSection}>
                <div className={styles.heading}>
                <h1>Welcome to Stidy</h1>
                </div>
                <div>
                    <div className={styles.inputSection}>
                    <input className={styles.input} type="text" placeholder="Username"/>
                    <input className={styles.input} type="text" placeholder="Password"/>
                    </div>
                    <div className={styles.btnSection}>
                    <button className={styles.btn1}>
                        <a href="">Log In</a>
                    </button>
                    <button onClick={() => {history.push("/signup")}} type="button" className={styles.btn2}>
                        <a href="">Sign Up</a>
                    </button>
                    </div>
                </div>
                

                </div>
                </div>
            </div>
        </div>
    );
}