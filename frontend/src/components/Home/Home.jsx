import styles from "./Home.module.scss";
import {useHistory} from "react-router";
import { gql, useMutation } from "@apollo/client";

const MUT = gql`mutation SignIn($name: String!, $pass: String!) { signIn(user: {name: $name, password: $pass})}`

export function Home (props) {
    let history = useHistory();
    let username = "";
    let password = "";
    const x = useMutation(MUT);

    if (x[1].loading) return 'Submitting...';
    if (x[1].error) return `Submission error! ${x[1].error.message}`;

    function signupClick() {
        history.push("/signup")
    }

    function handleUsername(node) {
        username = node;
    }

    function handlePassword(node) {
        password = node;
    }

    function login(e) {
        e.preventDefault();
        x[0]({variables: {name: username.value, pass: password.value}}).then(r => {
            props.sendJWT(r.data.toString());
            history.push('/mainpage')
        });
        username.value = '';
        password.value = '';
    }

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
                        <form onSubmit={login}>
                            <div className={styles.inputSection}>
                                <input className={styles.input} ref={handleUsername} type="text" placeholder="Username"/>
                                <input className={styles.input} ref={handlePassword} type="text" placeholder="Password"/>
                            </div>
                            <div className={styles.btnSection}>
                                <button type={"submit"} className={styles.btn1}>
                                    Log In
                                </button>
                                <button onClick={signupClick} type="button" className={styles.btn2}>
                                    Sign Up
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    );
}